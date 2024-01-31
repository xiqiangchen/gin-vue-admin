package dsp

import (
	dbid "github.com/flipped-aurora/gin-vue-admin/server/dsp/bid"
	"github.com/flipped-aurora/gin-vue-admin/server/dsp/bid/adapter"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp/bid"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"strings"
)

type BidService struct {
}

func (bidService *BidService) SendMsg(msg []byte) {
	utils.SendMsg(global.GVA_KAFKA_PRODUCER, global.GVA_CONFIG.Dsp.Bid.Topic, msg)
}

func (bidService *BidService) Bid(adxId int, body []byte) (resp []byte, offer bool) {
	// 从对接适配器中获取适配器
	adxAdapter := adapter.GetAdapter(adxId)
	if req, err := adxAdapter.From(body); err != nil {
		global.GVA_LOG.Error("协议转换失败", zap.Error(err))
		return nil, false
	} else if campaigns := filters(req); len(campaigns) == 0 {
		return nil, false
	}

	return
}

// 筛选符合条件的账户、计划、活动
func filters(req *bid.BidRequest) (campaigns []*ad.Campaign) {
	// 基础过滤
	// 计划状态过滤
	// 计划投放周期过滤
	// 计划预算过滤
	// 计划曝光频控过滤
	// 计划点击频控过滤
	// 活动状态过滤
	// 虚拟活动过滤
	// 活动投放周期过滤
	// 活动投放时段过滤
	// 活动预算过滤
	// 活动价格初步过滤
	// 活动报告频控
	// 活动点击频控
	// 活动黑白名单过滤
	// 活动定向包过滤
	// 模板创意素材过滤
	// 价格过滤
	// 计算出价
	// 填充创意
	// 填充曝光
	// 填充点击
	// 响应

	// 状态过滤、投放周期过滤、投放时间段、预算过滤都在dbid.ActiveCampaigns
	// 曝光频次过滤
	campaigns = filterByFrequencies(req, dbid.ActiveCampaigns)

	return
}

// 参与竞价，选出竞得者
func bids() {

}

// 补充必要信息
func fill() {

}

func filterByFrequencies(req *bid.BidRequest, cs []*ad.Campaign) (campaigns []*ad.Campaign) {
	for _, c := range cs {
		if filterByFrequency(req, c.GetImpFrequencyKey(), c.GetImpFrequencyMinute()) {
			continue
		}
		if filterByFrequency(req, c.GetClkFrequencyKey(), c.GetClkFrequencyMinute()) {
			continue
		}
	}
	return
}

func filterByFrequency(req *bid.BidRequest, frequencyKey, frequency int) bool {
	if v, exists := dbid.AdFrequency[frequencyKey]; !exists {
		return false
	} else if dev := req.GetDevice(); dev != nil {
		switch strings.ToLower(dev.GetOs()) {
		case "ios":
			if len(dev.GetCaid()) > 0 {
				if times, e := v.Get(dev.GetCaid()); e {
					if times.(int) >= frequency {
						return true
					}
				}
			}
		case "android":
			if len(dev.GetOaidmd5()) > 0 {
				if times, e := v.Get(dev.GetOaidmd5()); e {
					if times.(int) >= frequency {
						return true
					}
				}
			} else if len(dev.GetDidmd5()) > 0 {
				if times, e := v.Get(dev.GetDidmd5()); e {
					if times.(int) >= frequency {
						return true
					}
				}
			}
		}
	}
	return false
}
