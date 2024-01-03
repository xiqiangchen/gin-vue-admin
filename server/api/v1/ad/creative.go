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

type CreativeApi struct {
}

var creativeService = service.ServiceGroupApp.AdServiceGroup.CreativeService

// CreateCreative 创建创意表
// @Tags Creative
// @Summary 创建创意表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ad.Creative true "创建创意表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /creative/createCreative [post]
func (creativeApi *CreativeApi) CreateCreative(c *gin.Context) {
	var creative ad.Creative
	err := c.ShouldBindJSON(&creative)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	creative.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"PlanId":     {utils.NotEmpty()},
		"CampaignId": {utils.NotEmpty()},
		"MaterialId": {utils.NotEmpty()},
	}
	if err := utils.Verify(creative, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := creativeService.CreateCreative(&creative); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCreative 删除创意表
// @Tags Creative
// @Summary 删除创意表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ad.Creative true "删除创意表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /creative/deleteCreative [delete]
func (creativeApi *CreativeApi) DeleteCreative(c *gin.Context) {
	var creative ad.Creative
	err := c.ShouldBindJSON(&creative)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	creative.DeletedBy = utils.GetUserID(c)
	if err := creativeService.DeleteCreative(creative); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCreativeByIds 批量删除创意表
// @Tags Creative
// @Summary 批量删除创意表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除创意表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /creative/deleteCreativeByIds [delete]
func (creativeApi *CreativeApi) DeleteCreativeByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := creativeService.DeleteCreativeByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCreative 更新创意表
// @Tags Creative
// @Summary 更新创意表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ad.Creative true "更新创意表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /creative/updateCreative [put]
func (creativeApi *CreativeApi) UpdateCreative(c *gin.Context) {
	var creative ad.Creative
	err := c.ShouldBindJSON(&creative)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	creative.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"PlanId":     {utils.NotEmpty()},
		"CampaignId": {utils.NotEmpty()},
		"MaterialId": {utils.NotEmpty()},
	}
	if err := utils.Verify(creative, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := creativeService.UpdateCreative(creative); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCreative 用id查询创意表
// @Tags Creative
// @Summary 用id查询创意表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ad.Creative true "用id查询创意表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /creative/findCreative [get]
func (creativeApi *CreativeApi) FindCreative(c *gin.Context) {
	var creative ad.Creative
	err := c.ShouldBindQuery(&creative)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recreative, err := creativeService.GetCreative(creative.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recreative": recreative}, c)
	}
}

// GetCreativeList 分页获取创意表列表
// @Tags Creative
// @Summary 分页获取创意表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adReq.CreativeSearch true "分页获取创意表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /creative/getCreativeList [get]
func (creativeApi *CreativeApi) GetCreativeList(c *gin.Context) {
	var pageInfo adReq.CreativeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.CreatedBy = utils.GetUserID(c)
	if list, total, err := creativeService.GetCreativeInfoList(pageInfo); err != nil {
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
