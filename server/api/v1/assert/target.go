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

type TargetApi struct {
}

var targetService = service.ServiceGroupApp.AssertServiceGroup.TargetService

// CreateTarget 创建定向包
// @Tags Target
// @Summary 创建定向包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body assert.Target true "创建定向包"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /target/createTarget [post]
func (targetApi *TargetApi) CreateTarget(c *gin.Context) {
	var target assert.Target
	err := c.ShouldBindJSON(&target)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	target.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(target, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := targetService.CreateTarget(&target); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTarget 删除定向包
// @Tags Target
// @Summary 删除定向包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body assert.Target true "删除定向包"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /target/deleteTarget [delete]
func (targetApi *TargetApi) DeleteTarget(c *gin.Context) {
	var target assert.Target
	err := c.ShouldBindJSON(&target)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	target.DeletedBy = utils.GetUserID(c)
	if err := targetService.DeleteTarget(target); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTargetByIds 批量删除定向包
// @Tags Target
// @Summary 批量删除定向包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除定向包"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /target/deleteTargetByIds [delete]
func (targetApi *TargetApi) DeleteTargetByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := targetService.DeleteTargetByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTarget 更新定向包
// @Tags Target
// @Summary 更新定向包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body assert.Target true "更新定向包"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /target/updateTarget [put]
func (targetApi *TargetApi) UpdateTarget(c *gin.Context) {
	var target assert.Target
	err := c.ShouldBindJSON(&target)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	target.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(target, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := targetService.UpdateTarget(target); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTarget 用id查询定向包
// @Tags Target
// @Summary 用id查询定向包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query assert.Target true "用id查询定向包"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /target/findTarget [get]
func (targetApi *TargetApi) FindTarget(c *gin.Context) {
	var target assert.Target
	err := c.ShouldBindQuery(&target)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retarget, err := targetService.GetTarget(target.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retarget": retarget}, c)
	}
}

// GetTargetList 分页获取定向包列表
// @Tags Target
// @Summary 分页获取定向包列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query assertReq.TargetSearch true "分页获取定向包列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /target/getTargetList [get]
func (targetApi *TargetApi) GetTargetList(c *gin.Context) {
	var pageInfo assertReq.TargetSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.CreatedBy = utils.GetUserID(c)

	if list, total, err := targetService.GetTargetInfoList(pageInfo); err != nil {
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
