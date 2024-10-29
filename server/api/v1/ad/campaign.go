package ad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/dsp/bid"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	adReq "github.com/flipped-aurora/gin-vue-admin/server/model/ad/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math"
)

type CampaignApi struct {
}

var campaignService = service.ServiceGroupApp.AdServiceGroup.CampaignService

// CreateCampaign 创建活动
// @Tags Campaign
// @Summary 创建活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ad.Campaign true "创建活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /campaign/createCampaign [post]
func (campaignApi *CampaignApi) CreateCampaign(c *gin.Context) {
	var campaign ad.Campaign
	err := c.ShouldBindJSON(&campaign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"PlanId": {utils.NotEmpty()},
		"Name":   {utils.NotEmpty()},
	}
	if err := utils.Verify(campaign, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if plan, err := planService.GetPlan(campaign.PlanId); err != nil {
		global.GVA_LOG.Error("获取计划失败!", zap.Error(err))
		response.FailWithMessage("获取计划失败", c)
		return
	} else {
		campaign.Plan = &plan
	}

	campaign.Parse()
	campaign.CreatedBy = utils.GetUserID(c)
	if err := campaignService.CreateCampaign(&campaign); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		//response.OkWithMessage("创建成功", c)
		response.OkWithDetailed(gin.H{"recampaign": campaign}, "创建成功", c)
	}
}

// DeleteCampaign 删除活动
// @Tags Campaign
// @Summary 删除活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ad.Campaign true "删除活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /campaign/deleteCampaign [delete]
func (campaignApi *CampaignApi) DeleteCampaign(c *gin.Context) {
	var campaign ad.Campaign
	err := c.ShouldBindJSON(&campaign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	campaign.DeletedBy = utils.GetUserID(c)
	if err := campaignService.DeleteCampaign(campaign); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCampaignByIds 批量删除活动
// @Tags Campaign
// @Summary 批量删除活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /campaign/deleteCampaignByIds [delete]
func (campaignApi *CampaignApi) DeleteCampaignByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := campaignService.DeleteCampaignByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCampaign 更新活动
// @Tags Campaign
// @Summary 更新活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ad.Campaign true "更新活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /campaign/updateCampaign [put]
func (campaignApi *CampaignApi) UpdateCampaign(c *gin.Context) {
	var campaign ad.Campaign
	err := c.ShouldBindJSON(&campaign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if plan, err := planService.GetPlan(campaign.PlanId); err != nil {
		global.GVA_LOG.Error("获取计划失败!", zap.Error(err))
		response.FailWithMessage("获取计划失败", c)
		return
	} else {
		campaign.Plan = &plan
	}
	campaign.Parse()
	campaign.UpdatedBy = utils.GetUserID(c)
	if err := campaignService.UpdateCampaign(campaign); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCampaign 用id查询活动
// @Tags Campaign
// @Summary 用id查询活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ad.Campaign true "用id查询活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /campaign/findCampaign [get]
func (campaignApi *CampaignApi) FindCampaign(c *gin.Context) {
	var campaign ad.Campaign
	err := c.ShouldBindQuery(&campaign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recampaign, err := campaignService.GetCampaign(campaign.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		if v, ok := bid.BudgetControl.GetBudgetRecord(recampaign.GetBudgetKey()); ok {
			recampaign.TodayCost = math.Round(v.TotalUsage*100) / 100
			recampaign.TodayImpression = v.TotalImpressions
		}
		response.OkWithData(gin.H{"recampaign": recampaign}, c)
	}
}

// GetCampaignList 分页获取活动列表
// @Tags Campaign
// @Summary 分页获取活动列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adReq.CampaignSearch true "分页获取活动列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /campaign/getCampaignList [get]
func (campaignApi *CampaignApi) GetCampaignList(c *gin.Context) {
	var pageInfo adReq.CampaignSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.CreatedBy = utils.GetUserID(c)
	if list, total, err := campaignService.GetCampaignInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		for i := range list {
			if v, ok := bid.BudgetControl.GetBudgetRecord(list[i].GetBudgetKey()); ok {
				list[i].TodayCost = math.Round(v.TotalUsage*100) / 100
				list[i].TodayImpression = v.TotalImpressions
				list[i].TodayClick = v.TotalClicks
			}
		}
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
