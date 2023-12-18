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

type VideoApi struct {
}

var videoService = service.ServiceGroupApp.ResourceServiceGroup.VideoService

// CreateVideo 创建视频
// @Tags Video
// @Summary 创建视频
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resource.Video true "创建视频"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /video/createVideo [post]
func (videoApi *VideoApi) CreateVideo(c *gin.Context) {
	var video resource.Video
	err := c.ShouldBindJSON(&video)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	video.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"VideoUrl": {utils.NotEmpty()},
	}
	if err := utils.Verify(video, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := videoService.CreateVideo(&video); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteVideo 删除视频
// @Tags Video
// @Summary 删除视频
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resource.Video true "删除视频"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /video/deleteVideo [delete]
func (videoApi *VideoApi) DeleteVideo(c *gin.Context) {
	var video resource.Video
	err := c.ShouldBindJSON(&video)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	video.DeletedBy = utils.GetUserID(c)
	if err := videoService.DeleteVideo(video); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteVideoByIds 批量删除视频
// @Tags Video
// @Summary 批量删除视频
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除视频"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /video/deleteVideoByIds [delete]
func (videoApi *VideoApi) DeleteVideoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := videoService.DeleteVideoByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateVideo 更新视频
// @Tags Video
// @Summary 更新视频
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resource.Video true "更新视频"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /video/updateVideo [put]
func (videoApi *VideoApi) UpdateVideo(c *gin.Context) {
	var video resource.Video
	err := c.ShouldBindJSON(&video)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	video.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"VideoUrl": {utils.NotEmpty()},
	}
	if err := utils.Verify(video, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := videoService.UpdateVideo(video); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindVideo 用id查询视频
// @Tags Video
// @Summary 用id查询视频
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query resource.Video true "用id查询视频"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /video/findVideo [get]
func (videoApi *VideoApi) FindVideo(c *gin.Context) {
	var video resource.Video
	err := c.ShouldBindQuery(&video)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if revideo, err := videoService.GetVideo(video.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"revideo": revideo}, c)
	}
}

// GetVideoList 分页获取视频列表
// @Tags Video
// @Summary 分页获取视频列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query resourceReq.VideoSearch true "分页获取视频列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /video/getVideoList [get]
func (videoApi *VideoApi) GetVideoList(c *gin.Context) {
	var pageInfo resourceReq.VideoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := videoService.GetVideoInfoList(pageInfo); err != nil {
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
