package bid

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/constant"
	"github.com/flipped-aurora/gin-vue-admin/server/dsp/link"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp/bid"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp/budget"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp/strategy"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"sync"
	"time"
)

var ActiveCampaigns []*ad.Campaign
var Campaigns map[uint]*ad.Campaign

var AdFrequency = sync.Map{}           // 计划曝光频控
var BudgetControl budget.BudgetControl // 计划消耗
var ClickLimit *strategy.ClickLimit
var LinkSystemClient *link.Client

func Init() {
	if BudgetControl == nil {
		// redis模式还是本地模式
		if global.GVA_CONFIG.Dsp.UseRedis {
			BudgetControl = budget.NewRedisBudgetControl(1 * time.Minute)
		} else {
			BudgetControl = budget.NewLocalBudgetControl(1 * time.Minute)
		}
	}
	if ClickLimit == nil {
		// redis模式还是本地模式
		if global.GVA_CONFIG.Dsp.UseRedis {
			ClickLimit = strategy.NewClickLimit(7 * 24 * time.Hour)
		} else {
			// TODO
		}
	}
	if LinkSystemClient == nil {
		LinkSystemClient = link.NewClient(link.DefaultConfig())
	}
}

func ResetCampaignData() {
	BudgetControl.CleanToday()
}

// 定期扫描符合投放条件的活动
func Load() error {
	plans := loadPlans()
	var campaigns []*ad.Campaign
	for _, plan := range plans {
		campaigns = append(campaigns, plan.RealCampaigns...)
		if plan.GetImpFrequency() > 0 && plan.GetImpFrequencyMinute() > 0 {
			if _, exist := AdFrequency.Load(plan.GetImpFrequencyKey()); !exist {
				if dr, err := utils.ParseDuration(fmt.Sprintf("%vm", *plan.ImpFrequencyMinute)); err != nil {
					global.GVA_LOG.Error("格式化曝光频控失败:", zap.Error(err))
				} else {
					AdFrequency.Store(plan.GetImpFrequencyKey(), local_cache.NewCache(
						local_cache.SetDefaultExpire(dr),
					))
				}
			}

		}
		if plan.GetClkFrequency() > 0 && plan.GetClkFrequencyMinute() > 0 {
			if _, exist := AdFrequency.Load(plan.GetClkFrequencyKey()); !exist {
				if dr, err := utils.ParseDuration(fmt.Sprintf("%vm", *plan.ImpFrequencyMinute)); err != nil {
					global.GVA_LOG.Error("格式化点击频控失败:", zap.Error(err))
				} else {
					AdFrequency.Store(plan.GetClkFrequencyKey(), local_cache.NewCache(
						local_cache.SetDefaultExpire(dr),
					))
				}
			}
		}
	}
	ActiveCampaigns = campaigns
	cs := make(map[uint]*ad.Campaign, len(campaigns))

	for _, c := range campaigns {
		c.Init()
		cs[c.ID] = c
		if c.GetImpFrequency() > 0 && c.GetImpFrequencyMinute() > 0 {
			if _, exist := AdFrequency.Load(c.GetImpFrequencyKey()); !exist {
				if dr, err := utils.ParseDuration(fmt.Sprintf("%vm", *c.ImpFrequencyMinute)); err != nil {
					global.GVA_LOG.Error("格式化曝光频控失败:", zap.Error(err))
				} else {
					AdFrequency.Store(c.GetImpFrequencyKey(), local_cache.NewCache(
						local_cache.SetDefaultExpire(dr),
					))
				}
			}

		}
		if c.GetClkFrequency() > 0 && c.GetClkFrequencyMinute() > 0 {
			if _, exist := AdFrequency.Load(c.GetClkFrequencyKey()); !exist {
				if dr, err := utils.ParseDuration(fmt.Sprintf("%vm", *c.ImpFrequencyMinute)); err != nil {
					global.GVA_LOG.Error("格式化点击频控失败:", zap.Error(err))
				} else {
					AdFrequency.Store(c.GetClkFrequencyKey(), local_cache.NewCache(
						local_cache.SetDefaultExpire(dr),
					))
				}
			}
		}

		// 预算、曝光限制
		if c.GetBudgetDaily() > 0 || c.GetBudgetTotal() > 0 || c.GetImpTotal() > 0 || c.GetImpDaily() > 0 {
			key := c.GetBudgetKey()
			global.GVA_LOG.Info("预算和消耗情况", zap.Any(key, BudgetControl.Get(key)))
			BudgetControl.SetLimits(key, float64(c.GetBudgetDaily()), float64(c.GetBudgetTotal()), c.GetImpDaily(), c.GetImpTotal(), c.GetClkDaily(), c.GetClkTotal())
		}
	}
	Campaigns = cs
	return nil
}

func loadPlans() (plans []*bid.Plan) {
	var ps []*ad.Plan
	db := global.GVA_DB.Model(&ad.Plan{})
	db.Where("status = ? AND (filter is null or filter = ?) AND (start_at <= ? OR start_at is null) AND (end_at >= ? OR end_at is null)", constant.StatusOn, constant.Pass,
		time.Now(), time.Now())

	if err := db.Preload("Campaigns").Preload("Campaigns.Target").Preload("Campaigns.Creatives").Preload("Campaigns.Creatives.Material").Find(&ps).Error; err != nil {

	} else {
		for _, p := range ps {
			var realCampaigns, virtuallyCampaigns []*ad.Campaign

			for _, c := range p.Campaigns {
				if c.Status != nil && !*(c.Status) {
					continue
				}
				if c.Filter != nil && *(c.Filter) {
					continue
				}

				if !c.IsInDays() || !c.IsInHours() {
					continue
				}

				if c.IsVirtually != nil && *(c.IsVirtually) {
					virtuallyCampaigns = append(virtuallyCampaigns, c)
				} else {
					realCampaigns = append(realCampaigns, c)
				}
				c.Plan = p
			}
			plans = append(plans, &bid.Plan{
				Plan:               *p,
				RealCampaigns:      realCampaigns,
				VirtuallyCampaigns: virtuallyCampaigns,
			})
		}
	}

	return
}
