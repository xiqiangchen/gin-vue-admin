package ad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PlanRouter struct {
}

// InitPlanRouter 初始化 广告计划 路由信息
func (s *PlanRouter) InitPlanRouter(Router *gin.RouterGroup) {
	planRouter := Router.Group("plan").Use(middleware.OperationRecord())
	planRouterWithoutRecord := Router.Group("plan")
	var planApi = v1.ApiGroupApp.AdApiGroup.PlanApi
	{
		planRouter.POST("createPlan", planApi.CreatePlan)             // 新建广告计划
		planRouter.DELETE("deletePlan", planApi.DeletePlan)           // 删除广告计划
		planRouter.DELETE("deletePlanByIds", planApi.DeletePlanByIds) // 批量删除广告计划
		planRouter.PUT("updatePlan", planApi.UpdatePlan)              // 更新广告计划
	}
	{
		planRouterWithoutRecord.GET("findPlan", planApi.FindPlan)       // 根据ID获取广告计划
		planRouterWithoutRecord.GET("getPlanList", planApi.GetPlanList) // 获取广告计划列表
	}
}
