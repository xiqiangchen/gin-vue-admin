// 自动生成模板Material
package resource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 素材库 结构体  Material
type Material struct {
	global.GVA_MODEL
	Url       string `json:"url" form:"url" gorm:"column:url;comment:视频或图片链接;size:191;"`        //视频链接
	Type      int    `json:"type" form:"type" gorm:"column:type;comment:素材类型;size:10;"`         //素材类型
	Format    string `json:"format" form:"format" gorm:"column:format;comment:视频格式;size:19;"`   //格式
	Width     int    `json:"width" form:"width" gorm:"column:width;comment:宽;size:19;"`         //宽
	Height    int    `json:"height" form:"height" gorm:"column:height;comment:高;size:19;"`      //高
	Comment   string `json:"comment" form:"comment" gorm:"column:comment;comment:备注;size:191;"` //备注
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 素材库 Material自定义表名 materials
func (*Material) TableName() string {
	return "materials"
}

func (m *Material) GetAbsoluteUrl() string {
	return global.GVA_CONFIG.System.URLPrefix + m.Url
}
