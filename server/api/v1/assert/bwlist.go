package assert

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/assert"
	assertReq "github.com/flipped-aurora/gin-vue-admin/server/model/assert/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BlackWhiteListApi struct {
}

var bwlistService = service.ServiceGroupApp.AssertServiceGroup.BlackWhiteListService

// CreateBlackWhiteList 创建黑白名单
// @Tags BlackWhiteList
// @Summary 创建黑白名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body assert.BlackWhiteList true "创建黑白名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /bwlist/createBlackWhiteList [post]
func (bwlistApi *BlackWhiteListApi) CreateBlackWhiteList(c *gin.Context) {
	var bwlist assert.BlackWhiteList
	err := c.ShouldBindJSON(&bwlist)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	bwlist.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(bwlist, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := bwlistService.CreateBlackWhiteList(&bwlist); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBlackWhiteList 删除黑白名单
// @Tags BlackWhiteList
// @Summary 删除黑白名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body assert.BlackWhiteList true "删除黑白名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bwlist/deleteBlackWhiteList [delete]
func (bwlistApi *BlackWhiteListApi) DeleteBlackWhiteList(c *gin.Context) {
	var bwlist assert.BlackWhiteList
	err := c.ShouldBindJSON(&bwlist)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	bwlist.DeletedBy = utils.GetUserID(c)
	if err := bwlistService.DeleteBlackWhiteList(bwlist); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBlackWhiteListByIds 批量删除黑白名单
// @Tags BlackWhiteList
// @Summary 批量删除黑白名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除黑白名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bwlist/deleteBlackWhiteListByIds [delete]
func (bwlistApi *BlackWhiteListApi) DeleteBlackWhiteListByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := bwlistService.DeleteBlackWhiteListByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBlackWhiteList 更新黑白名单
// @Tags BlackWhiteList
// @Summary 更新黑白名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body assert.BlackWhiteList true "更新黑白名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bwlist/updateBlackWhiteList [put]
func (bwlistApi *BlackWhiteListApi) UpdateBlackWhiteList(c *gin.Context) {
	var bwlist assert.BlackWhiteList
	err := c.ShouldBindJSON(&bwlist)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	bwlist.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(bwlist, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := bwlistService.UpdateBlackWhiteList(bwlist); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBlackWhiteList 用id查询黑白名单
// @Tags BlackWhiteList
// @Summary 用id查询黑白名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query assert.BlackWhiteList true "用id查询黑白名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bwlist/findBlackWhiteList [get]
func (bwlistApi *BlackWhiteListApi) FindBlackWhiteList(c *gin.Context) {
	var bwlist assert.BlackWhiteList
	err := c.ShouldBindQuery(&bwlist)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rebwlist, err := bwlistService.GetBlackWhiteList(bwlist.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebwlist": rebwlist}, c)
	}
}

// GetBlackWhiteListList 分页获取黑白名单列表
// @Tags BlackWhiteList
// @Summary 分页获取黑白名单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query assertReq.BlackWhiteListSearch true "分页获取黑白名单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bwlist/getBlackWhiteListList [get]
func (bwlistApi *BlackWhiteListApi) GetBlackWhiteListList(c *gin.Context) {
	var pageInfo assertReq.BlackWhiteListSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := bwlistService.GetBlackWhiteListInfoList(pageInfo); err != nil {
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
