package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
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

	if err := clk.Check(); err != nil {
		response.IllegalWithMessage("非法请求", c)
		return
	}

	// 格式化
	clk.Parse()

	// 进入统计
	for _, cl := range clk.Expand() {
		clickService.SendMsg(cl.Marshal())
	}

}
