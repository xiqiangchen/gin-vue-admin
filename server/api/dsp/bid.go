package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

type BidApi struct {
}

var bidService = service.ServiceGroupApp.DspGroup.BidService

func (bidApi *BidApi) Req(c *gin.Context) {
	var imp dsp.Track
	if err := c.ShouldBindQuery(&imp); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := imp.Check(); err != nil {
		response.IllegalWithMessage("非法请求", c)
		return
	}

	// 进入统计

}
