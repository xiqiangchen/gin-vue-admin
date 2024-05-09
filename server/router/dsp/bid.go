package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/dsp"
	"github.com/gin-gonic/gin"
)

type BidRouter struct {
}

// InitBidRouter 初始化 广告计划 路由信息
func (s *BidRouter) InitBidRouter(Router *gin.RouterGroup) {
	bidRouter := Router.Group("bid")
	var bidApi = dsp.ApiGroupApp.BidApi
	{
		bidRouter.POST("rtb", bidApi.Rtb) // 更新广告计划
	}
}
