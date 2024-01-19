package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/dsp"
	"github.com/gin-gonic/gin"
)

type TrackRouter struct {
}

// InitClickRouter 初始化 监控 路由信息
func (s *TrackRouter) InitTrackRouter(Router *gin.RouterGroup) {
	trackRouter := Router.Group("track")
	var impressionApi = dsp.ApiGroupApp.ImpressionApi
	{
		trackRouter.GET("imp", impressionApi.ImpressionTrack) // 更新广告计划
	}
	var clickApi = dsp.ApiGroupApp.ClickApi
	{
		trackRouter.GET("clk", clickApi.ClickTrack) // 更新广告计划
	}
}
