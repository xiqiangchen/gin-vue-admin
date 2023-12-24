package resource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resource"
	resourceReq "github.com/flipped-aurora/gin-vue-admin/server/model/resource/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MaterialApi struct {
}

var materialService = service.ServiceGroupApp.ResourceServiceGroup.MaterialService

// CreateMaterial 创建素材库
// @Tags Material
// @Summary 创建素材库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resource.Material true "创建素材库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /material/createMaterial [post]
func (materialApi *MaterialApi) CreateMaterial(c *gin.Context) {
	var material resource.Material
	err := c.ShouldBindJSON(&material)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var typ int
	if len(material.VideoUrl) > 0 {
		typ = 2
	} else {
		typ = 1
	}
	material.Type = &typ

	material.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Type": {utils.NotEmpty()},
	}
	if err := utils.Verify(material, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := materialService.CreateMaterial(&material); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMaterial 删除素材库
// @Tags Material
// @Summary 删除素材库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resource.Material true "删除素材库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /material/deleteMaterial [delete]
func (materialApi *MaterialApi) DeleteMaterial(c *gin.Context) {
	var material resource.Material
	err := c.ShouldBindJSON(&material)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	material.DeletedBy = utils.GetUserID(c)
	if err := materialService.DeleteMaterial(material); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMaterialByIds 批量删除素材库
// @Tags Material
// @Summary 批量删除素材库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除素材库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /material/deleteMaterialByIds [delete]
func (materialApi *MaterialApi) DeleteMaterialByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := materialService.DeleteMaterialByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMaterial 更新素材库
// @Tags Material
// @Summary 更新素材库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resource.Material true "更新素材库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /material/updateMaterial [put]
func (materialApi *MaterialApi) UpdateMaterial(c *gin.Context) {
	var material resource.Material
	err := c.ShouldBindJSON(&material)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var typ int
	if len(material.VideoUrl) > 0 {
		typ = 2
	} else {
		typ = 1
	}
	material.Type = &typ

	material.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Type": {utils.NotEmpty()},
	}
	if err := utils.Verify(material, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := materialService.UpdateMaterial(material); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMaterial 用id查询素材库
// @Tags Material
// @Summary 用id查询素材库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query resource.Material true "用id查询素材库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /material/findMaterial [get]
func (materialApi *MaterialApi) FindMaterial(c *gin.Context) {
	var material resource.Material
	err := c.ShouldBindQuery(&material)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rematerial, err := materialService.GetMaterial(material.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rematerial": rematerial}, c)
	}
}

// GetMaterialList 分页获取素材库列表
// @Tags Material
// @Summary 分页获取素材库列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query resourceReq.MaterialSearch true "分页获取素材库列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /material/getMaterialList [get]
func (materialApi *MaterialApi) GetMaterialList(c *gin.Context) {
	var pageInfo resourceReq.MaterialSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := materialService.GetMaterialInfoList(pageInfo); err != nil {
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
