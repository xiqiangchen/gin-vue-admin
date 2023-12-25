// 自动生成模板Plan
package ad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 广告计划 结构体  Plan
type Plan struct {
	global.GVA_MODEL
	Name               string      `json:"name" form:"name" gorm:"column:name;comment:名称;size:191;"`                                                //名称
	Desc               string      `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:191;"`                                                //描述
	Status             *bool       `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`                                           //状态
	Mode               *int        `json:"mode" form:"mode" gorm:"column:mode;comment:投放方式;size:10;"`                                               //投放方式
	Timezone           *int        `json:"timezone" form:"timezone" gorm:"column:timezone;comment:时区:utc;size:10;"`                                 //时区:utc
	StartAt            *time.Time  `json:"startAt" form:"startAt" gorm:"column:start_at;comment:开始时间;"`                                             //开始时间
	EndAt              *time.Time  `json:"endAt" form:"endAt" gorm:"column:end_at;comment:结束时间;"`                                                   //结束时间
	BudgetTotal        *int        `json:"budgetTotal" form:"budgetTotal" gorm:"column:budget_total;comment:总预算,元;size:19;"`                        //总预算,千分
	BudgetDaily        *int        `json:"budgetDaily" form:"budgetDaily" gorm:"column:budget_daily;comment:每日预算,元;size:19;"`                       //每日预算,千分
	ImpTotal           *int        `json:"impTotal" form:"impTotal" gorm:"column:imp_total;comment:总曝光数;size:19;"`                                  //总曝光数
	ImpDaily           *int        `json:"impDaily" form:"impDaily" gorm:"column:imp_daily;comment:每日曝光数;size:19;"`                                 //每日曝光数
	ImpFrequency       *int        `json:"impFrequency" form:"impFrequency" gorm:"column:imp_frequency;comment:曝光频制;size:19;"`                      //曝光频制
	ImpFrequencyMinute *int        `json:"impFrequencyMinute" form:"impFrequencyMinute" gorm:"column:imp_frequency_minute;comment:曝光频控周期;size:19;"` //曝光频控周期
	ClkFrequency       *int        `json:"clkFrequency" form:"clkFrequency" gorm:"column:clk_frequency;comment:点击频控;size:19;"`                      //点击频控
	ClkFrequencyMinute *int        `json:"clkFrequencyMinute" form:"clkFrequencyMinute" gorm:"column:clk_frequency_minute;comment:点击频控周期;size:19;"` //点击频控周期
	CtrMax             *float64    `json:"ctrMax" form:"ctrMax" gorm:"column:ctr_max;comment:最小点击率，单位%;size:10;"`                                   //最小点击率，单位0.1%
	CtrMin             *float64    `json:"ctrMin" form:"ctrMin" gorm:"column:ctr_min;comment:最大点击率，单位%;size:10;"`                                   //最大点击率，单位0.1%
	CreatedBy          uint        `gorm:"column:created_by;comment:创建者"`
	UpdatedBy          uint        `gorm:"column:updated_by;comment:更新者"`
	DeletedBy          uint        `gorm:"column:deleted_by;comment:删除者"`
	Campaigns          []*Campaign `json:"campaigns" gorm:"foreignKey:plan_id"`
}

// TableName 广告计划 Plan自定义表名 plans
func (Plan) TableName() string {
	return "plans"
}
