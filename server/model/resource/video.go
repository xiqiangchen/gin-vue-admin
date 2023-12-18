// 自动生成模板Video
package resource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 视频 结构体  Video
type Video struct {
	global.GVA_MODEL
	VideoUrl  string `json:"video_url" form:"video_url" gorm:"column:video_url;comment:;"`        //视频内部链接
	ImageUrl  string `json:"image_url" form:"image_url" gorm:"column:image_url;comment:视频封面图链接;"` //视频封面图链接
	Format    string `json:"format" form:"format" gorm:"column:format;comment:视频格式;size:19;"`     //视频格式
	Width     *int   `json:"width" form:"width" gorm:"column:width;comment:宽;"`                   //宽
	Height    *int   `json:"height" form:"height" gorm:"column:height;comment:高;"`                //高
	Comment   string `json:"comment" form:"comment" gorm:"column:comment;comment:备注;"`            // 备注
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 视频 Video自定义表名 videos
func (Video) TableName() string {
	return "videos"
}
