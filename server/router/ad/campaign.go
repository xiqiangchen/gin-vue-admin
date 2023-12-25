package ad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CampaignRouter struct {
}

// InitCampaignRouter 初始化 活动 路由信息
func (s *CampaignRouter) InitCampaignRouter(Router *gin.RouterGroup) {
	campaignRouter := Router.Group("campaign").Use(middleware.OperationRecord())
	campaignRouterWithoutRecord := Router.Group("campaign")
	var campaignApi = v1.ApiGroupApp.AdApiGroup.CampaignApi
	{
		campaignRouter.POST("createCampaign", campaignApi.CreateCampaign)             // 新建活动
		campaignRouter.DELETE("deleteCampaign", campaignApi.DeleteCampaign)           // 删除活动
		campaignRouter.DELETE("deleteCampaignByIds", campaignApi.DeleteCampaignByIds) // 批量删除活动
		campaignRouter.PUT("updateCampaign", campaignApi.UpdateCampaign)              // 更新活动
	}
	{
		campaignRouterWithoutRecord.GET("findCampaign", campaignApi.FindCampaign)       // 根据ID获取活动
		campaignRouterWithoutRecord.GET("getCampaignList", campaignApi.GetCampaignList) // 获取活动列表
	}
}
