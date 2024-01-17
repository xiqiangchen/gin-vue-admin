package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/dsp"
	"github.com/gin-gonic/gin"
)

type ClickRouter struct {
}

// InitClickRouter 初始化 广告计划 路由信息
func (s *ClickRouter) InitClickRouter(Router *gin.RouterGroup) {
	clickRouter := Router.Group("click")
	var clickApi = dsp.ApiGroupApp.ClickApi
	{
		clickRouter.GET("clk", clickApi.ClickTrack) // 更新广告计划
	}
}
