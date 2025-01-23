// 自动生成模板Target
package assert

import (
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 定向包 结构体  Target
type Target struct {
	global.GVA_MODEL
	Name        string              `json:"name" form:"name" gorm:"column:name;comment:名称;"`                           // 名称
	Desc        string              `json:"desc" form:"desc" gorm:"column:desc;comment:描述;"`                           // 描述
	AdFormat    string              `json:"ad_format" form:"ad_format" gorm:"column:ad_format;comment:广告形式;size:255;"` // 广告形式
	AdFormats   map[int]struct{}    `json:"-" form:"-" gorm:"-"`
	DeviceType  string              `json:"device_type" form:"device_type" gorm:"column:device_type;comment:设备类型;size:255;"` // 设备类型
	DeviceTypes map[int]struct{}    `json:"-" form:"-" gorm:"-"`
	Os          string              `json:"os" form:"os" gorm:"column:os;comment:操作系统类型;size:255;"` // 操作系统
	Oses        map[int]struct{}    `json:"-" form:"-" gorm:"-"`
	TargetType  string              `json:"target_type" form:"target_type" gorm:"column:target_type;comment:定向类型;size:255;"` // 定向类型
	TargetTypes map[int]struct{}    `json:"-" form:"-" gorm:"-"`
	Region      string              `json:"region" form:"region" gorm:"column:region;comment:行政区域;size:255;"` // 行政区域
	Regions     map[string]struct{} `json:"-" form:"-" gorm:"-"`
	Adx         string              `json:"adx" form:"adx" gorm:"column:adx;comment:行政区域;size:255;"` // 行政区域
	Adxs        map[int]struct{}    `json:"-" form:"-" gorm:"-"`
	Gender      string              `json:"gender" form:"gender" gorm:"column:gender;comment:性别;size:255;"` // 性别
	Genders     map[int]struct{}    `json:"-" form:"-" gorm:"-"`
	CreatedBy   uint                `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint                `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint                `gorm:"column:deleted_by;comment:删除者"`
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

	if len(t.AdFormat) > 0 && t.AdFormats == nil {
		rs := strings.Split(t.AdFormat, ",")
		ts := make(map[int]struct{})
		for _, r := range rs {
			a, _ := strconv.Atoi(r)
			if a != 0 {
				ts[a] = struct{}{}
			}
		}
		t.AdFormats = ts
	}

	if len(t.DeviceType) > 0 && t.DeviceTypes == nil {
		rs := strings.Split(t.DeviceType, ",")
		ts := make(map[int]struct{})
		for _, r := range rs {
			a, _ := strconv.Atoi(r)
			if a != 0 {
				ts[a] = struct{}{}
			}
		}
		t.DeviceTypes = ts
	}

	if len(t.Os) > 0 && t.Oses == nil {
		rs := strings.Split(t.Os, ",")
		ts := make(map[int]struct{})
		for _, r := range rs {
			a, _ := strconv.Atoi(r)
			if a != 0 {
				ts[a] = struct{}{}
			}
		}
		t.Oses = ts
	}

	if len(t.TargetType) > 0 && t.TargetTypes == nil {
		rs := strings.Split(t.TargetType, ",")
		ts := make(map[int]struct{})
		for _, r := range rs {
			a, _ := strconv.Atoi(r)
			if a != 0 {
				ts[a] = struct{}{}
			}
		}
		t.TargetTypes = ts
	}

	if len(t.Gender) > 0 && t.Genders == nil {
		rs := strings.Split(t.Gender, ",")
		ts := make(map[int]struct{})
		for _, r := range rs {
			a, _ := strconv.Atoi(r)
			if a != 0 {
				ts[a] = struct{}{}
			}
		}
		t.Genders = ts
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
	if t.Oses != nil && len(t.Oses) > 0 {
		osValue := 0
		if strings.ToLower(os) == "ios" {
			osValue = 1
		} else if strings.ToLower(os) == "android" {
			osValue = 2
		}
		if _, ok := t.Oses[osValue]; ok {
			return true
		}
		return false
	}
	return true
}

func (t *Target) InPlatform(mobile int) bool {
	if t.DeviceTypes != nil && len(t.DeviceTypes) > 0 {
		if _, ok := t.DeviceTypes[mobile]; ok {
			return true
		}
		return false
	}
	return true
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

func (t *Target) InAdFormat(format int) bool {
	if t.AdFormats != nil && len(t.AdFormats) > 0 {
		if _, ok := t.AdFormats[format]; ok {
			return true
		}
		return false
	}
	return true
}

func (t *Target) InTargetType(targetType int) bool {
	if t.TargetTypes != nil && len(t.TargetTypes) > 0 {
		if _, ok := t.TargetTypes[targetType]; ok {
			return true
		}
		return false
	}
	return true
}

func (t *Target) InGender(gender int) bool {
	if t.Genders != nil && len(t.Genders) > 0 {
		if _, ok := t.Genders[gender]; ok {
			return true
		}
		return false
	}
	return true
}
