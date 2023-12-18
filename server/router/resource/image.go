package resource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ImageRouter struct {
}

// InitImageRouter 初始化 图片库 路由信息
func (s *ImageRouter) InitImageRouter(Router *gin.RouterGroup) {
	imageRouter := Router.Group("image").Use(middleware.OperationRecord())
	imageRouterWithoutRecord := Router.Group("image")
	var imageApi = v1.ApiGroupApp.ResourceApiGroup.ImageApi
	{
		imageRouter.POST("createImage", imageApi.CreateImage)   // 新建图片库
		imageRouter.DELETE("deleteImage", imageApi.DeleteImage) // 删除图片库
		imageRouter.DELETE("deleteImageByIds", imageApi.DeleteImageByIds) // 批量删除图片库
		imageRouter.PUT("updateImage", imageApi.UpdateImage)    // 更新图片库
	}
	{
		imageRouterWithoutRecord.GET("findImage", imageApi.FindImage)        // 根据ID获取图片库
		imageRouterWithoutRecord.GET("getImageList", imageApi.GetImageList)  // 获取图片库列表
	}
}
