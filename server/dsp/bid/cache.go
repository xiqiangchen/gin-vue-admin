package bid

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/constant"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp/bid"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"time"
)

var ActiveCampaigns []*ad.Campaign
var AdFrequency = make(map[int]local_cache.Cache) // 曝光、点击控制

func Load() error {
	plans := loadPlans()
	var campaigns []*ad.Campaign
	for _, plan := range plans {
		campaigns = append(campaigns, plan.RealCampaigns...)
		if plan.GetImpFrequency() > 0 && plan.GetImpFrequencyMinute() > 0 {
			if _, exist := AdFrequency[plan.GetImpFrequencyKey()]; !exist {
				if dr, err := utils.ParseDuration(fmt.Sprintf("%vm", *plan.ImpFrequencyMinute)); err != nil {
					fmt.Println("error:", err)
				} else {
					AdFrequency[plan.GetImpFrequencyKey()] = local_cache.NewCache(
						local_cache.SetDefaultExpire(dr),
					)
				}
			}

		}
		if plan.GetClkFrequency() > 0 && plan.GetClkFrequencyMinute() > 0 {
			if _, exist := AdFrequency[plan.GetClkFrequencyKey()]; !exist {
				if dr, err := utils.ParseDuration(fmt.Sprintf("%vm", *plan.ClkFrequencyMinute)); err != nil {
					fmt.Println("error:", err)
				} else {
					AdFrequency[plan.GetClkFrequencyKey()] = local_cache.NewCache(
						local_cache.SetDefaultExpire(dr),
					)
				}
			}
		}
	}
	ActiveCampaigns = campaigns

	for _, c := range campaigns {
		if c.GetImpFrequency() > 0 && c.GetImpFrequencyMinute() > 0 {
			if _, exist := AdFrequency[c.GetImpFrequencyKey()]; !exist {
				if dr, err := utils.ParseDuration(fmt.Sprintf("%vm", *c.ImpFrequencyMinute)); err != nil {
					fmt.Println("error:", err)
				} else {
					AdFrequency[c.GetImpFrequencyKey()] = local_cache.NewCache(
						local_cache.SetDefaultExpire(dr),
					)
				}
			}

		}
		if c.GetClkFrequency() > 0 && c.GetClkFrequencyMinute() > 0 {
			if _, exist := AdFrequency[c.GetClkFrequencyKey()]; !exist {
				if dr, err := utils.ParseDuration(fmt.Sprintf("%vm", *c.ClkFrequencyMinute)); err != nil {
					fmt.Println("error:", err)
				} else {
					AdFrequency[c.GetClkFrequencyKey()] = local_cache.NewCache(
						local_cache.SetDefaultExpire(dr),
					)
				}
			}
		}
	}

	return nil
}

func loadPlans() (plans []*bid.Plan) {
	var ps []*ad.Plan
	db := global.GVA_DB.Model(&ad.Plan{})
	db.Where("status = ? AND filter = ? AND start_at > ? AND end_at < ?", constant.StatusOn, constant.Pass,
		time.Now(), time.Now())

	if err := db.Preload("Campaigns").Find(&ps).Error; err != nil {

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

				if !c.IsInHours() {
					continue
				}

				if c.IsVirtually != nil && *(c.IsVirtually) {
					virtuallyCampaigns = append(virtuallyCampaigns, c)
				} else {
					realCampaigns = append(realCampaigns, c)
				}

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
