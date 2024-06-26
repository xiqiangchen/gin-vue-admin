package dsp

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"strconv"
	"strings"
)

type Track struct {
	AdxId       int    `json:"adx_id,omitempty" form:"ch,omitempty"`      // 渠道id
	SpotId      int    `json:"spot_id,omitempty" form:"sp,omitempty"`     // 广告位id
	TemplateId  int    `json:"template_id,omitempty" form:"tp,omitempty"` // 模板id
	UserId      int    `json:"user_id,omitempty" form:"u,omitempty"`      // 用户id
	PlanId      int    `json:"plan_id,omitempty" form:"p,omitempty"`      // 计划id
	CampaignId  int    `json:"campaign_id,omitempty" form:"c,omitempty"`  // 活动id
	CreativeId  int    `json:"creative_id,omitempty" form:"cr,omitempty"` // 创意id
	MaterialId  int    `json:"material_id,omitempty" form:"m,omitempty"`  // 素材id
	RequestId   string `json:"-" form:"rid,omitempty"`                    // 请求id
	ClickId     string `json:"-" form:"ckid,omitempty"`                   // 平台唯一id
	Device      int    `json:"device,omitempty" form:"dv,omitempty"`      // 设备
	Os          int    `json:"os,omitempty" form:"os,omitempty"`          // 系统
	App         string `json:"-" form:"ap,omitempty"`                     // 应用
	AppId       uint64 `json:"app,omitempty" form:"ap,omitempty"`         // 应用id
	Imei5       string `json:"-" form:"im5,omitempty"`                    // imei 的 md5 摘要，32位
	Idfa5       string `json:"-" form:"ifa5,omitempty"`                   // IOS 6+的设备id字段的md5，32位
	Caid        string `json:"-" form:"cid,omitempty"`                    // IOS 14后的设备id字段
	Caid5       string `json:"-" form:"cid5,omitempty"`                   // IOS 14后的设备id字段的md5，32位
	CaidVersion string `json:"-" form:"cidv,omitempty"`                   // caid版本号
	Oaid        string `json:"-" form:"oid,omitempty"`                    // Android Q及更高版本的设备号，32位
	Oaid5       string `json:"-" form:"oid5,omitempty"`                   // Android Q及更高版本的设备号的md5摘要，32位
	IP          string `json:"-" form:"ip,omitempty"`                     // IP地址
	UserAgent   string `json:"-" form:"ua,omitempty"`                     // user-agent
	ClickTs     int64  `json:"-" form:"ts,omitempty"`                     // 点击ts
	Sign        string `json:"-" form:"sign,omitempty"`                   // 验参
	MultiTrack  string `json:"-" form:"mul,omitempty"`                    // 多活动、创意、素材
	Metrics
}

type Impression struct {
	Track
}

type Click struct {
	Track
}

func (track *Track) Check() error {
	return nil
}

func (track *Track) Parse() {
	if len(track.App) > 0 {
		if appId, err := strconv.Atoi(track.App); err != nil {
			track.AppId = utils.HashCode(track.App)
		} else {
			track.AppId = uint64(appId)
		}
	}
}

func (track *Track) Clone() *Track {
	return &Track{
		Metrics: Metrics{
			Impression: track.Impression,
			Click:      track.Click,
		},
		AdxId:      track.AdxId,
		UserId:     track.UserId,
		PlanId:     track.PlanId,
		CampaignId: track.CampaignId,
		CreativeId: track.CreativeId,
		MaterialId: track.MaterialId,
		Os:         track.Os,
		Device:     track.Device,
		SpotId:     track.SpotId,
		TemplateId: track.TemplateId,
		AppId:      track.AppId,
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

func (imp *Impression) Check() error {
	return nil
}

func (clk *Click) Check() error {
	return nil
}

func (imp *Impression) Parse() error {
	imp.Impression = 1
	return nil
}

func (clk *Click) Parse() error {
	clk.Click = 1
	return nil
}
