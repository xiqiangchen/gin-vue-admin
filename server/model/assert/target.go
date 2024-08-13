// 自动生成模板Target
package assert

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 定向包 结构体  Target
type Target struct {
	global.GVA_MODEL
	Name       string `json:"name" form:"name" gorm:"column:name;comment:名称;"`                        // 名称
	Desc       string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;"`                        // 描述
	AdFormat   *int   `json:"ad_format" form:"ad_format" gorm:"column:ad_format;comment:广告形式;"`       // 广告形式
	DeviceType *int   `json:"device_type" form:"device_type" gorm:"column:device_type;comment:设备类型;"` // 设备类型
	Os         *int   `json:"os" form:"os" gorm:"column:os;comment:操作系统类型;"`                          // 操作系统
	TargetType *int   `json:"target_type" form:"target_type" gorm:"column:target_type;comment:定向类型;"` // 定向类型
	Region     string `json:"region" form:"region" gorm:"column:region;comment:行政区域;size:255;"`       // 行政区域
	Gender     *int   `json:"gender" form:"gender" gorm:"column:gender;comment:性别;"`                  // 性别
	CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 定向包 Target自定义表名 targets
func (Target) TableName() string {
	return "targets"
}
