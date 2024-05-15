package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/dsp/bid/adapter"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	bid_adapter "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/bid"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"strconv"
)

type BidApi struct {
}

var bidService = service.ServiceGroupApp.DspGroup.BidService

func (bidApi *BidApi) Rtb(c *gin.Context) {
	adx, ok := c.GetQuery("adx")
	if !ok {
		response.IllegalWithMessage("adx必填", c)
		return
	}
	adxId, _ := strconv.Atoi(adx)
	// 从对接适配器中获取适配器
	adxAdapter := adapter.GetAdapter(adxId)
	var req *bid_adapter.BidRequest
	var resp any

	if bodyBytes, err := io.ReadAll(c.Request.Body); err != nil {
		response.NoContent(c)
		global.GVA_LOG.Error("渠道解释失败", zap.Error(err))
	} else if req, err = adxAdapter.From(c.Request.Header, bodyBytes); err != nil {
		global.GVA_LOG.Error("协议转换失败", zap.Error(err))
		response.NoContent(c)
	} else if bresp, offer := bidService.Bid(req, c); !offer {
		response.NoContent(c)
	} else if resp, err = adxAdapter.To(bresp); err != nil {
		global.GVA_LOG.Error("bidResp转换协议失败", zap.Error(err))
	} else {
		response.AutoContent(resp, c)
	}
}
