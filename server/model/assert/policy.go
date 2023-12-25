// 自动生成模板Policy
package assert

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 出价策略 结构体  Policy
type Policy struct {
	global.GVA_MODEL
	Platform  *int     `json:"platform" form:"platform" gorm:"column:platform;comment:平台渠道;size:19;"`      //平台渠道
	Spot      *int     `json:"spot" form:"spot" gorm:"column:spot;comment:广告位;size:19;"`                   //广告位
	Publisher *int     `json:"publisher" form:"publisher" gorm:"column:publisher;comment:发布者;size:19;"`    //发布者
	Bundle    *int     `json:"bundle" form:"bundle" gorm:"column:bundle;comment:bundle;size:19;"`          //bundle
	Format    *int     `json:"format" form:"format" gorm:"column:format;comment:广告形式;size:10;"`            //广告形式
	Identity  *int     `json:"identity" form:"identity" gorm:"column:identity;comment:是否有设备id信息;size:10;"` //是否有设备id信息
	Os        *int     `json:"os" form:"os" gorm:"column:os;comment:操作系统;size:10;"`                        //操作系统
	Region    *int     `json:"region" form:"region" gorm:"column:region;comment:地区;size:19;"`              //地区
	Price     *float64 `json:"price" form:"price" gorm:"column:price;comment:平均出价;size:10;"`               //平均出价
	Scope     *int     `json:"scope" form:"scope" gorm:"column:scope;comment:浮动范围(%);size:10;"`            //浮动范围(%)
	CreatedBy uint     `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint     `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint     `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 出价策略 Policy自定义表名 policies
func (Policy) TableName() string {
	return "policies"
}
