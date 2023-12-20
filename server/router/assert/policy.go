package assert

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PolicyRouter struct {
}

// InitPolicyRouter 初始化 出价策略 路由信息
func (s *PolicyRouter) InitPolicyRouter(Router *gin.RouterGroup) {
	policyRouter := Router.Group("policy").Use(middleware.OperationRecord())
	policyRouterWithoutRecord := Router.Group("policy")
	var policyApi = v1.ApiGroupApp.AssertApiGroup.PolicyApi
	{
		policyRouter.POST("createPolicy", policyApi.CreatePolicy)   // 新建出价策略
		policyRouter.DELETE("deletePolicy", policyApi.DeletePolicy) // 删除出价策略
		policyRouter.DELETE("deletePolicyByIds", policyApi.DeletePolicyByIds) // 批量删除出价策略
		policyRouter.PUT("updatePolicy", policyApi.UpdatePolicy)    // 更新出价策略
	}
	{
		policyRouterWithoutRecord.GET("findPolicy", policyApi.FindPolicy)        // 根据ID获取出价策略
		policyRouterWithoutRecord.GET("getPolicyList", policyApi.GetPolicyList)  // 获取出价策略列表
	}
}
