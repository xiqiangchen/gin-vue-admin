package dsp

import (
	dbid "github.com/flipped-aurora/gin-vue-admin/server/dsp/bid"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type ImpressionApi struct {
}

var impressionService = service.ServiceGroupApp.DspGroup.ImpressionService

func (impressionApi *ImpressionApi) ImpressionTrack(c *gin.Context) {
	var imp dsp.Impression
	if err := c.ShouldBindQuery(&imp); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := imp.Validate(); err != nil {
		global.GVA_LOG.Warn("非法曝光：", zap.Any("query", c.Request.URL.RawQuery))
		response.IllegalWithMessage("非法请求", c)
		return
	}

	// 格式化
	imp.Parse()
	imp.Impression = 1

	// 进入统计
	for _, im := range imp.Expand() {
		//impressionService.SendMsg(im.Marshal())
		global.GVA_LOG.Info("收到曝光：", zap.ByteString("imp", im.Marshal()))
		dbid.BudgetControl.Update(im.GetCampaignBudgetKey(), im.RequestId, im.Price/1000, im.Impression, 0)
	}

	if len(imp.RedirectUrl) > 0 {
		c.Redirect(http.StatusFound, imp.RedirectUrl)
	} else {
		response.OkWithNoContent(c)
	}
}
