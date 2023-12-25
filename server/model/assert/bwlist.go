// 自动生成模板BlackWhiteList
package assert

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
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
	CreatedBy         uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy         uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy         uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 黑白名单 BlackWhiteList自定义表名 bwlist
func (BlackWhiteList) TableName() string {
	return "bwlist"
}
