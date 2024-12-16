package dsp

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/dsp/bid/pricer"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type Track struct {
	AdxId       int    `json:"adx_id,omitempty" form:"ch,omitempty"`        // 渠道id
	Spot        string `json:"spot,omitempty" form:"sp,omitempty"`          // 广告位id
	SpotId      uint64 `json:"spot_id,omitempty" form:"-"`                  // 广告位id
	TemplateId  int    `json:"template_id,omitempty" form:"tp,omitempty"`   // 模板id
	UserId      int    `json:"user_id,omitempty" form:"u,omitempty"`        // 用户id
	PlanId      int    `json:"plan_id,omitempty" form:"p,omitempty"`        // 计划id
	CampaignId  int    `json:"campaign_id,omitempty" form:"c,omitempty"`    // 活动id
	CreativeId  int    `json:"creative_id,omitempty" form:"cr,omitempty"`   // 创意id
	MaterialId  int    `json:"material_id,omitempty" form:"m,omitempty"`    // 素材id
	DeeplinkId  int    `json:"deeplink_id,omitempty" form:"dp,omitempty"`   // dpid
	DeeplinkId2 int    `json:"deeplink_id2,omitempty" form:"dp2,omitempty"` // 活动内deeplink行数
	AdType      int    `json:"ad_type,omitempty" form:"adt,omitempty"`      // 0: banner； 1: native； 2: video
	RequestId   string `json:"-" form:"rid,omitempty"`                      // 请求id
	ClickId     string `json:"-" form:"ckid,omitempty"`                     // 平台唯一id
	Device      int    `json:"device,omitempty" form:"dv,omitempty"`        // 设备
	Os          string `json:"-" form:"os,omitempty"`                       // 系统
	OsId        int    `json:"os,omitempty" form:"-"`                       // 系统
	App         string `json:"-" form:"ap,omitempty"`                       // 应用
	AppId       uint64 `json:"app,omitempty" form:"-"`                      // 应用id
	Publisher   string `json:"-" form:"puer,omitempty"`                     // 发布者
	PublisherId uint64 `json:"puer,omitempty" form:"-"`                     // 发布者id
	Site        string `json:"-" form:"site,omitempty"`                     // 网站
	SiteId      uint64 `json:"site,omitempty" form:"-"`                     // 网站id
	Did         string `json:"-" form:"did,omitempty"`                      // 设备id
	Did5        string `json:"-" form:"did5,omitempty"`                     // 设备id 的 md5 摘要，32位
	Imei5       string `json:"-" form:"im5,omitempty"`                      // imei 的 md5 摘要，32位
	Idfa5       string `json:"-" form:"ifa5,omitempty"`                     // IOS 6+的设备id字段的md5，32位
	Caid        string `json:"-" form:"cid,omitempty"`                      // IOS 14后的设备id字段
	Caid5       string `json:"-" form:"cid5,omitempty"`                     // IOS 14后的设备id字段的md5，32位
	CaidVersion string `json:"-" form:"cidv,omitempty"`                     // caid版本号
	Oaid        string `json:"-" form:"oid,omitempty"`                      // Android Q及更高版本的设备号，32位
	Oaid5       string `json:"-" form:"oid5,omitempty"`                     // Android Q及更高版本的设备号的md5摘要，32位
	IP          string `json:"-" form:"ip,omitempty"`                       // IP地址
	Country     string `json:"-" form:"cny,omitempty"`                      // 国家
	UserAgent   string `json:"-" form:"ua,omitempty"`                       // user-agent
	BidTs       int64  `json:"-" form:"ts,omitempty"`                       // 竞价ts
	Sign        string `json:"-" form:"sign,omitempty"`                     // 验参
	MultiTrack  string `json:"-" form:"mul,omitempty"`                      // 多活动、创意、素材
	RedirectUrl string `json:"-" form:"reurl,omitempty"`                    // 重定向地址
	Metrics
}

type Impression struct {
	Track
}

type Click struct {
	Track
}

func (track *Track) Validate() error {
	return nil
}

func (track *Track) Parse() {
	if len(track.App) > 0 {
		track.AppId = utils.GetId(track.App)
	}
	if len(track.Publisher) > 0 {
		track.PublisherId = utils.GetId(track.Publisher)
	}
	if len(track.Site) > 0 {
		track.SiteId = utils.GetId(track.Site)
	}
	if len(track.Spot) > 0 {
		track.SpotId = utils.GetId(track.Spot)
	}
	switch strings.ToUpper(track.Os) {
	case "android", "harmony", "harmonyos":
		track.OsId = 1
	case "ios":
		track.OsId = 2
	default:

	}
}

func (track *Track) GetCampaignBudgetKey() string {
	return "budget_" + strconv.Itoa(int(track.UserId*10000000+track.PlanId*10000+track.CampaignId*10)+1)
}

func (track *Track) GetPlanBudgetKey() string {
	return "budget_" + strconv.Itoa(int(track.UserId*10000+track.PlanId*10)+1)
}

func (track *Track) GetCampaignDeeplinkKey() string {
	return fmt.Sprintf("dp_%d_%d_%d", track.UserId, track.CampaignId, track.DeeplinkId2)
	//return "dp_" + strconv.Itoa(int(track.UserId*10000000+track.PlanId*10000+track.CampaignId*10)+1)
}

func (track *Track) Clone() *Track {
	return &Track{
		Metrics: Metrics{
			Impression: track.Impression,
			Click:      track.Click,
			Landing:    track.Landing,
			Price:      track.Price,
		},
		AdxId:       track.AdxId,
		UserId:      track.UserId,
		PlanId:      track.PlanId,
		CampaignId:  track.CampaignId,
		CreativeId:  track.CreativeId,
		MaterialId:  track.MaterialId,
		OsId:        track.OsId,
		Device:      track.Device,
		SpotId:      track.SpotId,
		TemplateId:  track.TemplateId,
		AppId:       track.AppId,
		PublisherId: track.PublisherId,
		SiteId:      track.SiteId,
	}
}

func (track *Track) Expand() (tracks []*Track) {
	tracks = append(tracks, track)
	if len(track.MultiTrack) > 0 {
		ts := strings.Split(track.MultiTrack, "-")
		for _, tr := range ts {
			t := track.Clone()
			ids := strings.Split(tr, "_")
			switch len(ids) {
			case 3:
				t.CampaignId, _ = strconv.Atoi(ids[0])
				t.CreativeId, _ = strconv.Atoi(ids[1])
				t.MaterialId, _ = strconv.Atoi(ids[2])
			case 4:
				t.PlanId, _ = strconv.Atoi(ids[0])
				t.CampaignId, _ = strconv.Atoi(ids[1])
				t.CreativeId, _ = strconv.Atoi(ids[2])
				t.MaterialId, _ = strconv.Atoi(ids[3])
			default:
				continue
			}
			tracks = append(tracks, t)
		}
	}
	return
}

func (track *Track) Marshal() []byte {
	byt, _ := json.Marshal(track)
	return byt
}

func (imp *Impression) Validate() error {
	var err error
	if imp.Price, err = pricer.DefaultPricer.Decode(imp.PriceSrc); err != nil {
		errors.New("价格解析失败: " + imp.PriceSrc)
	}
	// 曝光超过24小时算作弊
	if imp.BidTs <= 0 || imp.BidTs > 0 && time.Now().Unix()-imp.BidTs > 24*60*60 {
		return errors.New("曝光时间超过24小时")
	}
	return nil
}

func (clk *Click) Validate() error {
	return nil
}
