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

type ClickApi struct {
}

var clickService = service.ServiceGroupApp.DspGroup.ClickService

func (clickApi *ClickApi) ClickTrack(c *gin.Context) {
	var clk dsp.Click
	if err := c.ShouldBindQuery(&clk); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := clk.Validate(); err != nil {
		response.IllegalWithMessage("非法请求", c)
		return
	}

	// 格式化
	clk.Parse()
	clk.Click = 1

	// 进入统计
	for _, cl := range clk.Expand() {
		global.GVA_LOG.Info("收到点击：", zap.ByteString("clk", cl.Marshal()))
		dbid.BudgetControl.Update(cl.GetCampaignBudgetKey(), cl.RequestId, 0, 0, cl.Click)

		//clickService.SendMsg(cl.Marshal())
	}

	if len(clk.RedirectUrl) > 0 {
		c.Redirect(http.StatusFound, clk.RedirectUrl)
	} else {
		response.OkWithNoContent(c)
	}
}
