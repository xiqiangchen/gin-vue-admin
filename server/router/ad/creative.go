package ad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CreativeRouter struct {
}

// InitCreativeRouter 初始化 创意表 路由信息
func (s *CreativeRouter) InitCreativeRouter(Router *gin.RouterGroup) {
	creativeRouter := Router.Group("creative").Use(middleware.OperationRecord())
	creativeRouterWithoutRecord := Router.Group("creative")
	var creativeApi = v1.ApiGroupApp.AdApiGroup.CreativeApi
	{
		creativeRouter.POST("createCreative", creativeApi.CreateCreative)             // 新建创意表
		creativeRouter.DELETE("deleteCreative", creativeApi.DeleteCreative)           // 删除创意表
		creativeRouter.DELETE("deleteCreativeByIds", creativeApi.DeleteCreativeByIds) // 批量删除创意表
		creativeRouter.PUT("updateCreative", creativeApi.UpdateCreative)              // 更新创意表
	}
	{
		creativeRouterWithoutRecord.GET("findCreative", creativeApi.FindCreative)       // 根据ID获取创意表
		creativeRouterWithoutRecord.GET("getCreativeList", creativeApi.GetCreativeList) // 获取创意表列表
	}
}
