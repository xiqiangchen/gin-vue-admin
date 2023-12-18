package resource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VideoRouter struct {
}

// InitVideoRouter 初始化 视频 路由信息
func (s *VideoRouter) InitVideoRouter(Router *gin.RouterGroup) {
	videoRouter := Router.Group("video").Use(middleware.OperationRecord())
	videoRouterWithoutRecord := Router.Group("video")
	var videoApi = v1.ApiGroupApp.ResourceApiGroup.VideoApi
	{
		videoRouter.POST("createVideo", videoApi.CreateVideo)             // 新建视频
		videoRouter.DELETE("deleteVideo", videoApi.DeleteVideo)           // 删除视频
		videoRouter.DELETE("deleteVideoByIds", videoApi.DeleteVideoByIds) // 批量删除视频
		videoRouter.PUT("updateVideo", videoApi.UpdateVideo)              // 更新视频
	}
	{
		videoRouterWithoutRecord.GET("findVideo", videoApi.FindVideo)       // 根据ID获取视频
		videoRouterWithoutRecord.GET("getVideoList", videoApi.GetVideoList) // 获取视频列表
	}
}
