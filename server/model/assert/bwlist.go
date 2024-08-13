// 自动生成模板BlackWhiteList
package assert

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"strings"
)

// 黑白名单 结构体  BlackWhiteList
type BlackWhiteList struct {
	global.GVA_MODEL
	Name              string `json:"name" form:"name" gorm:"column:name;comment:名称;size:191;"`                                      //名称
	Desc              string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:255;"`                                      //描述
	PlatformWhitelist string `json:"platformWhitelist" form:"platformWhitelist" gorm:"column:platform_whitelist;comment:平台、渠道白名单;"` //平台、渠道白名单
	PlatformBlacklist string `json:"platformBlacklist" form:"platformBlacklist" gorm:"column:platform_blacklist;comment:平台、渠道黑名单;"` //平台、渠道黑名单
	SpotWhitelist     string `json:"spotWhitelist" form:"spotWhitelist" gorm:"column:spot_whitelist;comment:广告位白名单;"`               //广告位白名单
	SpotBlacklist     string `json:"spotBlacklist" form:"spotBlacklist" gorm:"column:spot_blacklist;comment:广告位黑名单;"`               //广告位黑名单
	BundleWhitelist   string `json:"bundleWhitelist" form:"bundleWhitelist" gorm:"column:bundle_whitelist;comment:应用白名单;"`          //应用白名单
	BundleBlacklist   string `json:"bundleBlacklist" form:"bundleBlacklist" gorm:"column:bundle_blacklist;comment:应用黑名单;"`          //应用黑名单
	SiteWhitelist     string `json:"siteWhitelist" form:"siteWhitelist" gorm:"column:site_whitelist;comment:网站域名白名单;"`              //网站域名白名单
	SiteBlacklist     string `json:"siteBlacklist" form:"siteBlacklist" gorm:"column:site_blacklist;comment:网站域名黑名单;"`              //网站域名黑名单
	DeviceWhitelist   string `json:"deviceWhitelist form:deviceWhitelist" gorm:"column:device_whitelist;comment:设备白名单;type:text;"`  // 设备白名单
	DeviceBlacklist   string `json:"deviceBlacklist form:deviceBlacklist" gorm:"column:device_blacklist;comment:设备黑名单;type:text;"`  // 设备黑名单

	deivceWLMap map[string]map[string]struct{} // 设备白名单
	deivceBLMap map[string]map[string]struct{} // 设备黑名单
	CreatedBy   uint                           `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint                           `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint                           `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 黑白名单 BlackWhiteList自定义表名 bwlist
func (BlackWhiteList) TableName() string {
	return "bwlist"
}

func (bw *BlackWhiteList) Init() {
	// 先弄设备黑白名单
	if len(bw.DeviceWhitelist) > 0 && bw.deivceWLMap == nil {
		bw.deivceWLMap = make(map[string]map[string]struct{})
		for _, line := range strings.Split(bw.DeviceWhitelist, "\n") {
			if ids := strings.Split(line, ","); len(ids) == 2 {
				if _, e := bw.deivceWLMap[ids[0]]; !e {
					bw.deivceWLMap[ids[0]] = make(map[string]struct{})
				}
				bw.deivceWLMap[ids[0]][ids[1]] = struct{}{}
			}
		}
	}
	if len(bw.DeviceBlacklist) > 0 && bw.deivceBLMap == nil {
		bw.deivceBLMap = make(map[string]map[string]struct{})
		for _, line := range strings.Split(bw.DeviceBlacklist, "\n") {
			if ids := strings.Split(line, ","); len(ids) == 2 {
				if _, e := bw.deivceBLMap[ids[0]]; !e {
					bw.deivceBLMap[ids[0]] = make(map[string]struct{})
				}
				bw.deivceBLMap[ids[0]][ids[1]] = struct{}{}
			}
		}
	}
}

func (bw *BlackWhiteList) HasDeviceWhileBlackList() bool {
	if len(bw.deivceWLMap) > 0 || len(bw.deivceBLMap) > 0 {
		return true
	}
	return false
}

func (bw *BlackWhiteList) IsDeviceWhileList(deviceType, deviceId string) bool {
	return bw.isWBList(bw.deivceWLMap, deviceType, deviceId)
}

func (bw *BlackWhiteList) IsDeviceBlackList(deviceType, deviceId string) bool {
	return bw.isWBList(bw.deivceBLMap, deviceType, deviceId)
}

func (bw *BlackWhiteList) isWBList(wblist map[string]map[string]struct{}, deviceType, deviceId string) bool {
	if wblist == nil {
		return false
	} else if v, e := wblist[deviceType]; e && v != nil {
		if _, e2 := v[deviceId]; e2 {
			return true
		} else {
			return false
		}
	}
	return false
}
