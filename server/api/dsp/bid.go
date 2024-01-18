package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BidApi struct {
}

var bidService = service.ServiceGroupApp.DspGroup.BidService

func (bidApi *BidApi) Req(c *gin.Context) {
	adx, ok := c.GetQuery("adx")
	if !ok {
		response.IllegalWithMessage("adx必填", c)
	}
	_, _ = strconv.Atoi(adx)

	// 从对接适配器中获取适配器

}
