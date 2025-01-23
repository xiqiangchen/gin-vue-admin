package openrtb_v2_6_1

type ExtTracks struct {
	ImpressionTracks []string `json:"impression_tracking_url,omitempty"`    // 曝光监测
	ClickTracks      []string `json:"click_tracking_urls,omitempty"`        // 点击监测
	StartAppTracks   []string `json:"start_app_urls,omitempty"`             // 唤醒app监测
	DownloadTracks   []string `json:"download_urls,omitempty"`              // 下载app
	DownloadUrl      string   `json:"download_url,omitempty"`               // 下载链接
	LandingUrl       string   `json:"ad_choices_destination_url,omitempty"` // 落地页，可以是302重定向的
	Deeplink         string   `json:"deeplink,omitempty"`                   // deeplink
	UniversalLink    string   `json:"universal_link,omitempty"`             // ios ulink
	FallbackUrl      string   `json:"fallback_url,omitempty"`               // fallback_url
	FallbackLogoUrl  string   `json:"fallback_logo,omitempty"`              // fallback_logo_url
	FallbackTitle    string   `json:"fallback_title,omitempty"`             // fallback_title
	BillingId        int64    `json:"billing_id,omitempty"`                 // billing id
	// 国内下载类必填
	// 安卓：app.name、app.bundle、app.ver、app.keywords、app.publisher.name、app.privacypolicy=1、app.ext.privacy_policy_url、app.ext.permissions
	// IOS：app.id、app.name、app.storeurl
	AppInfo *App `json:"app_info,omitempty"`
}
