package assert

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TargetRouter struct {
}

// InitTargetRouter 初始化 定向包 路由信息
func (s *TargetRouter) InitTargetRouter(Router *gin.RouterGroup) {
	targetRouter := Router.Group("target").Use(middleware.OperationRecord())
	targetRouterWithoutRecord := Router.Group("target")
	var targetApi = v1.ApiGroupApp.AssertApiGroup.TargetApi
	{
		targetRouter.POST("createTarget", targetApi.CreateTarget)             // 新建定向包
		targetRouter.DELETE("deleteTarget", targetApi.DeleteTarget)           // 删除定向包
		targetRouter.DELETE("deleteTargetByIds", targetApi.DeleteTargetByIds) // 批量删除定向包
		targetRouter.PUT("updateTarget", targetApi.UpdateTarget)              // 更新定向包
	}
	{
		targetRouterWithoutRecord.GET("findTarget", targetApi.FindTarget)       // 根据ID获取定向包
		targetRouterWithoutRecord.GET("getTargetList", targetApi.GetTargetList) // 获取定向包列表
	}
}
