package dsp

type Track struct {
	Type        int    `json:"ty,omitempty" form:"ty,omitempty"`     // 监测类型，默认0曝光，1点击
	ChannelId   int    `json:"ch,omitempty" form:"ch,omitempty"`     // 渠道id
	UserId      int    `json:"u,omitempty" form:"u,omitempty"`       // 用户id
	PlanId      int    `json:"p,omitempty" form:"p,omitempty"`       // 计划id
	CampaignId  int    `json:"c,omitempty" form:"c,omitempty"`       // 活动id
	CreativeId  int    `json:"cr,omitempty" form:"cr,omitempty"`     // 创意id
	MaterialId  int    `json:"m,omitempty" form:"m,omitempty"`       // 素材id
	RequestId   string `form:"rid,omitempty"`                        // 请求id
	ClickId     string `json:"ckid,omitempty" form:"ckid,omitempty"` // 平台唯一id
	Os          int    `json:"os,omitempty" form:"os,omitempty"`     // 系统
	Imei5       string `json:"im5,omitempty" form:"im5,omitempty"`   // imei 的 md5 摘要，32位
	Idfa5       string `json:"ifa5,omitempty" form:"ifa5,omitempty"` // IOS 6+的设备id字段的md5，32位
	Caid        string `json:"cid,omitempty" form:"cid,omitempty"`   // IOS 14后的设备id字段
	Caid5       string `json:"cid5,omitempty" form:"cid5,omitempty"` // IOS 14后的设备id字段的md5，32位
	CaidVersion string `json:"cidv,omitempty" form:"cidv,omitempty"` // caid版本号
	Oaid        string `json:"oid,omitempty" form:"oid,omitempty"`   // Android Q及更高版本的设备号，32位
	Oaid5       string `json:"oid5,omitempty" form:"oid5,omitempty"` // Android Q及更高版本的设备号的md5摘要，32位
	IP          string `json:"ip,omitempty" form:"ip,omitempty"`     // IP地址
	UserAgent   string `json:"ua,omitempty" form:"ua,omitempty"`     // user-agent
	ClickTs     int64  `json:"ts,omitempty" form:"ts,omitempty"`     // 点击ts
	Sign        string `json:"sign,omitempty" form:"sign,omitempty"` // 验参
}

func (track *Track) Check() error {
	return nil
}
