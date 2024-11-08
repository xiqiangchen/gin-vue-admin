// 自动生成模板Target
package assert

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"strconv"
	"strings"
)

// 定向包 结构体  Target
type Target struct {
	global.GVA_MODEL
	Name       string              `json:"name" form:"name" gorm:"column:name;comment:名称;"`                        // 名称
	Desc       string              `json:"desc" form:"desc" gorm:"column:desc;comment:描述;"`                        // 描述
	AdFormat   *int                `json:"ad_format" form:"ad_format" gorm:"column:ad_format;comment:广告形式;"`       // 广告形式
	DeviceType *int                `json:"device_type" form:"device_type" gorm:"column:device_type;comment:设备类型;"` // 设备类型
	Os         *int                `json:"os" form:"os" gorm:"column:os;comment:操作系统类型;"`                          // 操作系统
	TargetType *int                `json:"target_type" form:"target_type" gorm:"column:target_type;comment:定向类型;"` // 定向类型
	Region     string              `json:"region" form:"region" gorm:"column:region;comment:行政区域;size:255;"`       // 行政区域
	Regions    map[string]struct{} `json:"-" form:"-" gorm:"-"`
	Adx        string              `json:"adx" form:"adx" gorm:"column:adx;comment:行政区域;size:255;"` // 行政区域
	Adxs       map[int]struct{}    `json:"-" form:"-" gorm:"-"`
	Gender     *int                `json:"gender" form:"gender" gorm:"column:gender;comment:性别;"` // 性别
	CreatedBy  uint                `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint                `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint                `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 定向包 Target自定义表名 targets
func (Target) TableName() string {
	return "targets"
}

func (t *Target) Init() {
	if len(t.Region) > 0 && t.Regions == nil {
		rs := strings.Split(t.Region, ",")
		ts := make(map[string]struct{})
		for _, r := range rs {
			if r != "0" {
				ts[r] = struct{}{}
			}
		}
		t.Regions = ts
	}
	if len(t.Adx) > 0 && t.Adxs == nil {
		rs := strings.Split(t.Adx, ",")
		ts := make(map[int]struct{})
		for _, r := range rs {
			a, _ := strconv.Atoi(r)
			if a != 0 {
				ts[a] = struct{}{}
			}
		}
		t.Adxs = ts
	}
}

func (t *Target) InRegion(region string) bool {
	if t.Regions != nil && len(t.Regions) > 0 && len(region) > 0 {
		if _, ok := t.Regions[region]; ok {
			return true
		} else {
			return false
		}
	}
	// 默认是
	return true
}

func (t *Target) InOs(os string) bool {
	if t.Os != nil && *t.Os > 0 {
		if *t.Os == 1 && strings.ToLower(os) == "ios" {
			return true
		} else if *t.Os == 2 && strings.ToLower(os) == "android" {
			return true
		}
		return false
	}
	return true
}

func (t *Target) InPlatform(mobile int) bool {
	if t.DeviceType != nil && *t.DeviceType > 0 {
		if *t.DeviceType == 1 && mobile == 1 {
			return true
		}
	}
	return false
}

func (t *Target) InAdx(adx int) bool {
	if t.Adxs != nil && len(t.Adxs) > 0 {
		if _, ok := t.Adxs[adx]; ok {
			return true
		} else {
			return false
		}
	}
	// 默认是
	return true
}
