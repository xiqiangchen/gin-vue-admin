package resource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MaterialRouter struct {
}

// InitMaterialRouter 初始化 素材库 路由信息
func (s *MaterialRouter) InitMaterialRouter(Router *gin.RouterGroup) {
	materialRouter := Router.Group("material").Use(middleware.OperationRecord())
	materialRouterWithoutRecord := Router.Group("material")
	var materialApi = v1.ApiGroupApp.ResourceApiGroup.MaterialApi
	{
		materialRouter.POST("createMaterial", materialApi.CreateMaterial)             // 新建素材库
		materialRouter.DELETE("deleteMaterial", materialApi.DeleteMaterial)           // 删除素材库
		materialRouter.DELETE("deleteMaterialByIds", materialApi.DeleteMaterialByIds) // 批量删除素材库
		materialRouter.PUT("updateMaterial", materialApi.UpdateMaterial)              // 更新素材库
	}
	{
		materialRouterWithoutRecord.GET("findMaterial", materialApi.FindMaterial)       // 根据ID获取素材库
		materialRouterWithoutRecord.GET("getMaterialList", materialApi.GetMaterialList) // 获取素材库列表
	}
}
