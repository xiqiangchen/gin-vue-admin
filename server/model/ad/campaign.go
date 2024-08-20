// 自动生成模板Campaign
package ad

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/assert"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// 活动 结构体  Campaign
type Campaign struct {
	global.GVA_MODEL
	PlanId             uint                   `json:"plan_id" form:"plan_id" gorm:"column:plan_id;comment:计划id"`                                                   // 关联标记
	Plan               *Plan                  `json:"plan"`                                                                                                        // 广告计划
	Name               string                 `json:"name" form:"name" gorm:"column:name;comment:名称;size:191;"`                                                    //名称
	Desc               string                 `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:191;"`                                                    //描述
	Status             *bool                  `json:"status" form:"status" gorm:"column:status;comment:状态;"`                                                       //描述
	Filter             *bool                  `json:"filter" form:"filter" gorm:"column:filter;comment:投放过滤;"`                                                     //状态
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
	Hours              *int64                 `json:"hours" form:"hours" gorm:"column:hours;comment:投放时间段，10位表示的二进制;size:30;"`                                     //投放时间段，10位表示的二进制
	TargetId           *int                   `json:"target_id" form:"target_id" gorm:"column:target_id;comment:定向包id;size:19;"`                                   //定向包id
	Target             *assert.Target         `json:"target" gorm:"foreignKey:target_id"`                                                                          // 定向包
	BwlistId           *int                   `json:"bwlist_id" form:"bwlist_id" gorm:"column:bwlist_id;comment:黑白名单id;size:19;"`                                  // 黑白名单
	BlackWhiteList     *assert.BlackWhiteList `json:"black_white_list" gorm:"foreignKey:bwlist_id"`                                                                //黑白名单id
	PolicyId           *int                   `json:"policy_id" form:"policy_id" gorm:"column:policy_id;comment:出价策略id;size:19;"`                                  //出价策略id
	Policy             *assert.Policy         `json:"policy" gorm:"foreignKey:policy_id"`                                                                          // 出价策略
	BidMethod          *int                   `json:"bid_method" form:"bid_method" gorm:"column:bid_method;comment:出价方式;size:10;"`                                 //出价方式
	BidPrice           *float64               `json:"bid_price" form:"bid_price" gorm:"column:bid_price;comment:出价策略;size:10;"`                                    //出价策略
	BidMode            *int                   `json:"bid_mode" form:"bid_mode" gorm:"column:bid_mode;comment:出价模式;"`                                               //出价模式
	BidRate            *float64               `json:"bid_rate" form:"bid_rate" gorm:"column:bid_rate;comment:指定出价率;"`                                              //指定出价率
	Brand              string                 `json:"brand" form:"brand" gorm:"column:brand;comment:品牌名称;size:191;"`                                               //品牌名称
	AllowVirtually     *bool                  `json:"allow_virtually" form:"allow_virtually" gorm:"column:allow_virtually;comment:允许虚拟;"`                          //允许虚拟
	IsVirtually        *bool                  `json:"is_virtually" form:"is_virtually" gorm:"column:is_virtually;comment:是否虚拟活动;"`                                 //允许虚拟
	CreativeMode       *int                   `json:"creative_mode" form:"creative_mode" gorm:"column:creative_mode;comment:创意方式;size:10;"`                        //创意方式
	ImpTrackUrl        string                 `json:"imp_track_url" form:"imp_track_url" gorm:"column:imp_track_url;comment:曝光监测;size:2048;"`                      //曝光监测
	ClickTrackUrl      string                 `json:"click_track_url" form:"click_track_url" gorm:"column:click_track_url;comment:点击监测;size:2048;"`                //点击监测
	H5                 string                 `json:"h5" form:"h5" gorm:"column:h5;comment:落地页;size:2048;"`                                                        //落地页
	Deeplink           string                 `json:"deeplink" form:"deeplink" gorm:"column:deeplink;comment:;size:2048;"`                                         //deeplink字段
	UniversalLink      string                 `json:"universal_link" form:"universal_link" gorm:"column:universal_link;comment:;size:2048;"`                       //universalLink字段
	Adm                string                 `json:"adm" form:"adm" gorm:"column:adm;comment:动态代码;type:text;"`                                                    //动态代码
	CreatedBy          uint                   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy          uint                   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy          uint                   `gorm:"column:deleted_by;comment:删除者"`
	Creatives          []*Creative            `json:"creatives" gorm:"foreignKey:campaign_id"`
	Images             map[int][]*Creative    `json:"images" gorm:"-"`
	Videos             map[int][]*Creative    `json:"Videos" gorm:"-"`
}

// TableName 活动 Campaign自定义表名 campaigns
func (*Campaign) TableName() string {
	return "campaigns"
}

func (c *Campaign) GetHours() []int {
	if c.Hours != nil {
		return utils.ToArrFromBitInt(*(c.Hours))
	}
	return nil
}

func (c *Campaign) IsInDays() bool {
	// "status = ? AND (filter is null or filter = ?) AND (start_at <= ? OR start_at is null) AND (end_at >= ? OR end_at is null)"
	return c.GetStartAt().Before(time.Now()) && c.GetEndAt().After(time.Now())
}

func (c *Campaign) IsInHours() bool {
	if c.Hours == nil || *c.Hours == 0 {
		return true
	}
	now := time.Now().Hour()

	hours := utils.IntToBinaryArray(int64(*c.Hours), 24)

	if hours[now] > 0 {
		return true
	}

	return false
}

func (c *Campaign) GetBidPrice() float64 {
	if c.BidPrice == nil {
		return 0
	}
	return *c.BidPrice
}

func (c *Campaign) GetBudgetKey() string {
	return "budget_" + strconv.Itoa(c.GetImpFrequencyKey())
}

func (c *Campaign) GetImpFrequencyKey() int {
	return int(c.CreatedBy*10000000+c.PlanId*10000+c.ID*10) + 1
}

func (c *Campaign) GetClkFrequencyKey() int {
	return int(c.CreatedBy*10000000+c.PlanId*10000+c.ID*10) + 2
}

func (c *Campaign) GetClkFrequency() int {
	if c.ClkFrequency != nil {
		return *c.ClkFrequency
	}
	return 0
}

func (c *Campaign) GetImpFrequency() int {
	if c.ImpFrequency != nil {
		return *c.ImpFrequency
	}
	return 0
}

func (c *Campaign) GetImpFrequencyMinute() int {
	if c.ImpFrequencyMinute != nil {
		return *c.ImpFrequencyMinute
	}
	return 0
}

func (c *Campaign) GetClkFrequencyMinute() int {
	if c.ClkFrequencyMinute != nil {
		return *c.ClkFrequencyMinute
	}
	return 0
}

func (c *Campaign) GetStatus() bool {
	if c.Status != nil {
		return *c.Status
	}
	return false
}
func (c *Campaign) GetFilter() bool {
	if c.Filter != nil {
		return *c.Filter
	}
	return false
}
func (c *Campaign) GetStartAt() time.Time {
	if c.StartAt != nil {
		return *c.StartAt
	}
	return time.Unix(1715149854, 0)
}

func (c *Campaign) GetEndAt() time.Time {
	if c.EndAt != nil {
		return *c.EndAt
	}
	return time.Unix(4081300254, 0)
}
func (c *Campaign) GetBudgetTotal() int {
	if c.BudgetTotal != nil {
		return *c.BudgetTotal
	}
	return math.MaxInt
}
func (c *Campaign) GetBudgetDaily() int {
	if c.BudgetDaily != nil {
		return *c.BudgetDaily
	}
	return math.MaxInt
}
func (c *Campaign) GetImpTotal() int {
	if c.ImpTotal != nil {
		return *c.ImpTotal
	}
	return math.MaxInt
}
func (c *Campaign) GetImpDaily() int {
	if c.ImpDaily != nil {
		return *c.ImpDaily
	}
	return math.MaxInt
}
func (c *Campaign) GetCtrMax() float64 {
	if c.CtrMax != nil {
		return *c.CtrMax
	}
	return math.MaxInt
}
func (c *Campaign) GetCtrMin() float64 {
	if c.CtrMin != nil {
		return *c.CtrMin
	}
	return 0
}
func (c *Campaign) GetTargetId() int {
	if c.TargetId != nil {
		return *c.TargetId
	}
	return 0
}
func (c *Campaign) GetBwlistId() int {
	if c.BwlistId != nil {
		return *c.BwlistId
	}
	return 0
}
func (c *Campaign) GetPolicyId() int {
	if c.PolicyId != nil {
		return *c.PolicyId
	}
	return 0
}
func (c Campaign) GetBidMethod() int {
	if c.BidMethod != nil {
		return *c.BidMethod
	}
	return 0
}
func (c *Campaign) GetBidMode() int {
	if c.BidMode != nil {
		return *c.BidMode
	}
	return 0
}
func (c *Campaign) GetAllowVirtually() bool {
	if c.AllowVirtually != nil {
		return *c.AllowVirtually
	}
	return false
}
func (c *Campaign) GetIsVirtually() bool {
	if c.IsVirtually != nil {
		return *c.IsVirtually
	}
	return false
}
func (c *Campaign) GetCreativeMode() int {
	if c.CreativeMode != nil {
		return *c.CreativeMode
	}
	return 0
}
func (c *Campaign) Parse() {
	if c.StartAt != nil {
		start := utils.SetDateStart(c.GetStartAt(), *c.Plan.Timezone)
		c.StartAt = &start
	}
	if c.EndAt != nil {
		end := utils.SetDateEnd(c.GetEndAt(), *c.Plan.Timezone)
		c.EndAt = &end
	}
}

func (c *Campaign) Init() {
	c.BuildCreatives()
	if c.BlackWhiteList != nil {
		c.BlackWhiteList.Init()
	}
	if c.Target != nil {
		c.Target.Init()
	}
}
func (c *Campaign) InRegion(region string) bool {
	if c.Target != nil {
		return c.Target.InRegion(region)
	}
	return true
}

func (c *Campaign) BuildCreatives() {
	c.Images = make(map[int][]*Creative, len(c.Creatives))
	c.Videos = make(map[int][]*Creative, len(c.Creatives))
	for _, cr := range c.Creatives {
		if m := cr.Material; m != nil {
			switch m.Type {
			case 1:
				c.Images[m.Width*10000+m.Height] = append(c.Images[m.Width*10000+m.Height], cr)
			case 2:
				c.Videos[m.Width*10000+m.Height] = append(c.Videos[m.Width*10000+m.Height], cr)
			}
		}
	}

}
func (c *Campaign) SelectCreative(creativeType, width, height int) (*Creative, bool) {
	var cMap map[int][]*Creative
	switch creativeType {
	case 1:
		cMap = c.Images
	case 2:
		cMap = c.Videos
	}
	if len(cMap) > 0 {
		if cr, exist := cMap[width*10000+height]; exist && len(cr) > 0 {
			return cr[rand.Intn(len(cr))], true
		}
		return c.getNearlyCreative(cMap, width, height)
	} else {
		return nil, false
	}
}

func (c *Campaign) getNearlyCreative(cmap map[int][]*Creative, w, h int) (*Creative, bool) {
	var minC *Creative
	var minRate = math.MaxFloat64
	rate := float64(w) / float64(h)
	for k, v := range cmap {
		r := math.Abs(rate - float64(k/10000)/float64(k%10000))
		if r < minRate {
			minRate = r
			minC = v[rand.Intn(len(v))]
		}
	}
	return minC, minC != nil
}

func (c *Campaign) GetBidRate() float64 {
	if c.BidRate != nil && *c.BidRate > 0 && *c.BidRate < 100 {
		return *c.BidRate
	}
	return 100
}

func (c *Campaign) GetAdm() string {
	return c.Adm
}

func (c *Campaign) buildAdmFor(imp, clk string) string {
	var imps, clks []string
	if len(c.ClickTrackUrl) > 0 {
		clks = append(clks, c.ClickTrackUrl)
	}
	if len(clk) > 0 {
		clks = append(clks, clk)
	}
	if len(c.ImpTrackUrl) > 0 {
		imps = append(imps, c.ImpTrackUrl)
	}
	if len(imp) > 0 {
		imps = append(imps, imp)
	}

	var adm string
	if len(clks) > 0 {
		admClk := ""
		for i, ck := range clks {
			admClk = admClk + fmt.Sprintf("const img%d = new Image(); img%d.src = '%s';", i, i, ck)
		}
		adm = `<script> document.addEventListener('click', function(event) { ` + admClk + `}); </script>`
	}
	if len(imps) > 0 {
		admImp := ""
		for _, im := range imps {
			admImp = admImp + fmt.Sprintf("<img style='width:1px;height:1px;position:absolute;bottom:0;right:0;' src='%s'>", im)
		}
		adm = admImp + adm

	}
	return adm
}

func (c *Campaign) BuildAdmForIframe(imp, clk string) string {
	iframe := "<iframe src='" + c.Adm + "' frameborder='no' scrolling='no' style='width: 100%; height: 100%'  allow='cross-origin-isolated'></iframe>"
	return c.buildAdmFor(imp, clk) + iframe
}

func (c *Campaign) BuildAdmForCode(imp, clk string) string {
	return c.buildAdmFor(imp, clk) + c.Adm
}

func (c *Campaign) FillTrackParams(params map[string]string) {
	params["u"] = strconv.Itoa(int(c.CreatedBy))
	params["p"] = strconv.Itoa(int(c.PlanId))
	params["c"] = strconv.Itoa(int(c.ID))
}
