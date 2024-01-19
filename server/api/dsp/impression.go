package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
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
	if err := imp.Check(); err != nil {
		response.IllegalWithMessage("非法请求", c)
		return
	}

	// 格式化
	imp.Parse()

	// 进入统计
	for _, im := range imp.Expand() {
		impressionService.SendMsg(im.Marshal())
	}
}
