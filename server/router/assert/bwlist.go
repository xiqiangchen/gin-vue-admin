package assert

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BlackWhiteListRouter struct {
}

// InitBlackWhiteListRouter 初始化 黑白名单 路由信息
func (s *BlackWhiteListRouter) InitBlackWhiteListRouter(Router *gin.RouterGroup) {
	bwlistRouter := Router.Group("bwlist").Use(middleware.OperationRecord())
	bwlistRouterWithoutRecord := Router.Group("bwlist")
	var bwlistApi = v1.ApiGroupApp.AssertApiGroup.BlackWhiteListApi
	{
		bwlistRouter.POST("createBlackWhiteList", bwlistApi.CreateBlackWhiteList)             // 新建黑白名单
		bwlistRouter.DELETE("deleteBlackWhiteList", bwlistApi.DeleteBlackWhiteList)           // 删除黑白名单
		bwlistRouter.DELETE("deleteBlackWhiteListByIds", bwlistApi.DeleteBlackWhiteListByIds) // 批量删除黑白名单
		bwlistRouter.PUT("updateBlackWhiteList", bwlistApi.UpdateBlackWhiteList)              // 更新黑白名单
	}
	{
		bwlistRouterWithoutRecord.GET("findBlackWhiteList", bwlistApi.FindBlackWhiteList)       // 根据ID获取黑白名单
		bwlistRouterWithoutRecord.GET("getBlackWhiteListList", bwlistApi.GetBlackWhiteListList) // 获取黑白名单列表
	}
}
