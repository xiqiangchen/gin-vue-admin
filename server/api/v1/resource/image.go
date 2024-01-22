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

type ImageApi struct {
}

var imageService = service.ServiceGroupApp.ResourceServiceGroup.ImageService

// CreateImage 创建图片库
// @Tags Image
// @Summary 创建图片库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resource.Image true "创建图片库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /image/createImage [post]
func (imageApi *ImageApi) CreateImage(c *gin.Context) {
	var image resource.Image
	err := c.ShouldBindJSON(&image)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	image.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Url": {utils.NotEmpty()},
	}
	if err := utils.Verify(image, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := imageService.CreateImage(&image); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteImage 删除图片库
// @Tags Image
// @Summary 删除图片库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resource.Image true "删除图片库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /image/deleteImage [delete]
func (imageApi *ImageApi) DeleteImage(c *gin.Context) {
	var image resource.Image
	err := c.ShouldBindJSON(&image)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	image.DeletedBy = utils.GetUserID(c)
	if err := imageService.DeleteImage(image); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteImageByIds 批量删除图片库
// @Tags Image
// @Summary 批量删除图片库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除图片库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /image/deleteImageByIds [delete]
func (imageApi *ImageApi) DeleteImageByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := imageService.DeleteImageByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateImage 更新图片库
// @Tags Image
// @Summary 更新图片库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resource.Image true "更新图片库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /image/updateImage [put]
func (imageApi *ImageApi) UpdateImage(c *gin.Context) {
	var image resource.Image
	err := c.ShouldBindJSON(&image)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	image.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Url": {utils.NotEmpty()},
	}
	if err := utils.Verify(image, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := imageService.UpdateImage(image); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindImage 用id查询图片库
// @Tags Image
// @Summary 用id查询图片库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query resource.Image true "用id查询图片库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /image/findImage [get]
func (imageApi *ImageApi) FindImage(c *gin.Context) {
	var image resource.Image
	err := c.ShouldBindQuery(&image)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reimage, err := imageService.GetImage(image.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reimage": reimage}, c)
	}
}

// GetImageList 分页获取图片库列表
// @Tags Image
// @Summary 分页获取图片库列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query resourceReq.ImageSearch true "分页获取图片库列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /image/getImageList [get]
func (imageApi *ImageApi) GetImageList(c *gin.Context) {
	var pageInfo resourceReq.ImageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := imageService.GetImageInfoList(pageInfo); err != nil {
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
