// 自动生成模板Creative
package ad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resource"
)

// 创意表 结构体  Creative
type Creative struct {
	global.GVA_MODEL
	PlanId     int                `json:"plan_id" form:"plan_id" gorm:"column:plan_id;comment:;size:19;"`             //planId字段
	CampaignId int                `json:"campaign_id" form:"campaign_id" gorm:"column:campaign_id;comment:;size:19;"` //campaignId字段
	MaterialId int                `json:"material_id" form:"material_id" gorm:"column:material_id;comment:;size:19;"` //materialId字段
	Status     *bool              `json:"status" form:"status" gorm:"column:status;comment:状态;"`                      //描述
	Title      string             `json:"title" form:"title" gorm:"column:title;comment:;size:50;"`                   //title字段
	Desc       string             `json:"desc" form:"desc" gorm:"column:desc;comment:;size:255;"`                     //desc字段
	Button     string             `json:"button" form:"button" gorm:"column:button;comment:;size:30;"`                //button字段
	CreatedBy  uint               `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint               `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint               `gorm:"column:deleted_by;comment:删除者"`
	Material   *resource.Material `json:"material"`
	Plan       Plan               `json:"plan"`
	Campaign   Campaign           `json:"campaign"`
}

// TableName 创意表 Creative自定义表名 creatives
func (*Creative) TableName() string {
	return "creatives"
}
