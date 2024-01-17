package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/dsp"
	"github.com/gin-gonic/gin"
)

type ImpressionRouter struct {
}

// InitImpressionRouter 初始化 广告计划 路由信息
func (s *ImpressionRouter) InitImpressionRouter(Router *gin.RouterGroup) {
	impressionRouter := Router.Group("impression")
	var impressionApi = dsp.ApiGroupApp.ImpressionApi
	{
		impressionRouter.GET("imp", impressionApi.ImpressionTrack) // 更新广告计划
	}
}
