package ad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	adReq "github.com/flipped-aurora/gin-vue-admin/server/model/ad/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PlanApi struct {
}

var planService = service.ServiceGroupApp.AdServiceGroup.PlanService

// CreatePlan 创建广告计划
// @Tags Plan
// @Summary 创建广告计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ad.Plan true "创建广告计划"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /plan/createPlan [post]
func (planApi *PlanApi) CreatePlan(c *gin.Context) {
	var plan ad.Plan
	err := c.ShouldBindJSON(&plan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	plan.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Status":   {utils.NotEmpty()},
		"Mode":     {utils.NotEmpty()},
		"Timezone": {utils.NotEmpty()},
		"StartAt":  {utils.NotEmpty()},
	}
	if err := utils.Verify(plan, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := planService.CreatePlan(&plan); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePlan 删除广告计划
// @Tags Plan
// @Summary 删除广告计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ad.Plan true "删除广告计划"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /plan/deletePlan [delete]
func (planApi *PlanApi) DeletePlan(c *gin.Context) {
	var plan ad.Plan
	err := c.ShouldBindJSON(&plan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	plan.DeletedBy = utils.GetUserID(c)
	if err := planService.DeletePlan(plan); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePlanByIds 批量删除广告计划
// @Tags Plan
// @Summary 批量删除广告计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除广告计划"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /plan/deletePlanByIds [delete]
func (planApi *PlanApi) DeletePlanByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := planService.DeletePlanByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePlan 更新广告计划
// @Tags Plan
// @Summary 更新广告计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ad.Plan true "更新广告计划"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /plan/updatePlan [put]
func (planApi *PlanApi) UpdatePlan(c *gin.Context) {
	var plan ad.Plan
	err := c.ShouldBindJSON(&plan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	plan.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Status":   {utils.NotEmpty()},
		"Mode":     {utils.NotEmpty()},
		"Timezone": {utils.NotEmpty()},
		"StartAt":  {utils.NotEmpty()},
	}
	if err := utils.Verify(plan, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := planService.UpdatePlan(plan); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPlan 用id查询广告计划
// @Tags Plan
// @Summary 用id查询广告计划
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ad.Plan true "用id查询广告计划"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /plan/findPlan [get]
func (planApi *PlanApi) FindPlan(c *gin.Context) {
	var plan ad.Plan
	err := c.ShouldBindQuery(&plan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if replan, err := planService.GetPlan(plan.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"replan": replan}, c)
	}
}

// GetPlanList 分页获取广告计划列表
// @Tags Plan
// @Summary 分页获取广告计划列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adReq.PlanSearch true "分页获取广告计划列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /plan/getPlanList [get]
func (planApi *PlanApi) GetPlanList(c *gin.Context) {
	var pageInfo adReq.PlanSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.CreatedBy = utils.GetUserID(c)
	if list, total, err := planService.GetPlanInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
