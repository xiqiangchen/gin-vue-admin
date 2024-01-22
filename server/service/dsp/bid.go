package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/dsp/bid/adapter"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp/bid"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
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

	return
}

// 参与竞价，选出竞得者
func bids() {

}

// 补充必要信息
func fill() {

}
