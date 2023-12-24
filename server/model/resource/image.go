// 自动生成模板Image
package resource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 图片库 结构体  Image
type Image struct {
	global.GVA_MODEL
	Url       string `json:"url" form:"url" gorm:"column:url;comment:链接;"`                     //链接
	Format    string `json:"format" form:"format" gorm:"column:format;comment:图片格式;size:19;"`  //图片格式
	Width     *int   `json:"width" form:"width" gorm:"column:width;comment:宽;"`                //宽
	Height    *int   `json:"height" form:"height" gorm:"column:height;comment:高;"`             //高
	Comment   string `json:"comment" form:"comment" gorm:"column:comment;comment:备注;size:30;"` //备注
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 图片库 Image自定义表名 images
func (Image) TableName() string {
	return "images"
}
