// 自动生成模板Campaign
package ad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/assert"
	"time"
)

// 活动 结构体  Campaign
type Campaign struct {
	global.GVA_MODEL
	PlanID             uint                   `json:"plan_id" form:"plan_id" gorm:"column:plan_id;comment:计划id"` // 关联标记
	Plan               Plan                   `json:"plan"`
	Name               string                 `json:"name" form:"name" gorm:"column:name;comment:名称;size:191;"`                                                    //名称
	Desc               string                 `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:191;"`                                                    //描述
	Status             *bool                  `json:"status" form:"status" gorm:"column:status;comment:状态;"`                                                       //状态
	StartAt            *time.Time             `json:"start_at" form:"start_at" gorm:"column:start_at;comment:开始时间;"`                                               //开始时间
	EndAt              *time.Time             `json:"end_at" form:"end_at" gorm:"column:end_at;comment:结束时间;"`                                                     //结束时间
	BudgetTotal        *int                   `json:"budget_total" form:"budget_total" gorm:"column:budget_total;comment:总预算,元;size:20;"`                          //总预算,元
	BudgetDaily        *int                   `json:"budget_daily" form:"budget_daily" gorm:"column:budget_daily;comment:每日预算,元;size:20;"`                         //每日预算,元
	ImpTotal           *int                   `json:"imp_total" form:"imp_total" gorm:"column:imp_total;comment:总曝光数;size:20;"`                                    //总曝光数
	ImpDaily           *int                   `json:"imp_daily" form:"imp_daily" gorm:"column:imp_daily;comment:每日曝光数;size:20;"`                                   //每日曝光数
	ImpFrequency       *int                   `json:"imp_frequency" form:"imp_frequency" gorm:"column:imp_frequency;comment:曝光频制;size:10;"`                        //曝光频制
	ImpFrequencyMinute *int                   `json:"imp_frequency_minute" form:"imp_frequency_minute" gorm:"column:imp_frequency_minute;comment:曝光频控周期;size:10;"` //曝光频控周期
	ClkFrequency       *int                   `json:"clk_frequency" form:"clk_frequency" gorm:"column:clk_frequency;comment:点击频控;size:10;"`                        //点击频控
	ClkFrequencyMinute *int                   `json:"clk_frequency_minute" form:"clk_frequency_minute" gorm:"column:clk_frequency_minute;comment:点击频控周期;size:10;"` //点击频控周期
	CtrMax             *float64               `json:"ctr_max" form:"ctr_max" gorm:"column:ctr_max;comment:最小点击率，单位%;"`                                             //最小点击率，单位%
	CtrMin             *float64               `json:"ctr_min" form:"ctr_min" gorm:"column:ctr_min;comment:最大点击率，单位%;"`                                             //最大点击率，单位%
	Hours              *int                   `json:"hours" form:"hours" gorm:"column:hours;comment:投放时间段，10位表示的二进制;size:10;"`                                     //投放时间段，10位表示的二进制
	TargetId           *int                   `json:"target_id" form:"target_id" gorm:"column:target_id;comment:定向包id;size:19;"`                                   //定向包id
	Target             *assert.Target         `json:"target" gorm:"foreignKey:target_id"`                                                                          // 定向包
	BwlistId           *int                   `json:"bwlist_id" form:"bwlist_id" gorm:"column:bwlist_id;comment:黑白名单id;size:19;"`                                  // 黑白名单
	BlackWhiteList     *assert.BlackWhiteList `json:"black_white_list" gorm:"foreignKey:bwlist_id"`                                                                //黑白名单id
	PolicyId           *int                   `json:"policy_id" form:"policy_id" gorm:"column:policy_id;comment:出价策略id;size:19;"`                                  //出价策略id
	Policy             *assert.Policy         `json:"policy" gorm:"foreignKey:policy_id"`                                                                          // 出价策略
	BidMethod          *int                   `json:"bid_method" form:"bid_method" gorm:"column:bid_method;comment:出价方式;size:10;"`                                 //出价方式
	BidPrice           *float64               `json:"bid_price" form:"bid_price" gorm:"column:bid_price;comment:出价策略;size:10;"`                                    //出价策略
	BidMode            *int                   `json:"bid_mode" form:"bid_mode" gorm:"column:bid_mode;comment:出价模式;"`                                               //出价模式
	Brand              string                 `json:"brand" form:"brand" gorm:"column:brand;comment:品牌名称;size:191;"`                                               //品牌名称
	AllowVirtually     *bool                  `json:"allow_virtually" form:"allow_virtually" gorm:"column:allow_virtually;comment:允许虚拟;"`                          //允许虚拟
	IsVirtually        *bool                  `json:"is_virtually" form:"is_virtually" gorm:"column:is_virtually;comment:是否虚拟活动;"`                                 //允许虚拟
	CreativeMode       *int                   `json:"creative_mode" form:"creative_mode" gorm:"column:creative_mode;comment:创意方式;size:10;"`                        //创意方式
	ImpTrackUrl        string                 `json:"imp_track_url" form:"imp_track_url" gorm:"column:imp_track_url;comment:曝光监测;size:2048;"`                      //曝光监测
	ClickTrackUrl      string                 `json:"click_track_url" form:"click_track_url" gorm:"column:click_track_url;comment:点击监测;size:2048;"`                //点击监测
	H5                 string                 `json:"h5" form:"h5" gorm:"column:h5;comment:落地页;size:2048;"`                                                        //落地页
	Deeplink           string                 `json:"deeplink" form:"deeplink" gorm:"column:deeplink;comment:;size:2048;"`                                         //deeplink字段
	UniversalLink      string                 `json:"universal_link" form:"universal_link" gorm:"column:universal_link;comment:;size:2048;"`                       //universalLink字段
	CreatedBy          uint                   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy          uint                   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy          uint                   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 活动 Campaign自定义表名 campaigns
func (Campaign) TableName() string {
	return "campaigns"
}
