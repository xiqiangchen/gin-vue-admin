package openrtb_v2_6

type ExtTracks struct {
	ImpressionTracks []string `json:"imptracks,omitempty"`     // 曝光监测
	ClickTracks      []string `json:"clktracks,omitempty"`     // 点击检查
	LandingUrl       string   `json:"landingurl,omitempty"`    // 落地页，可以是302重定向的
	Deeplink         string   `json:"deeplink,omitempty"`      // deeplink
	UniversalLink    string   `json:"universallink,omitempty"` // ios ulink
	FallbackUrl      string   `json:"fallbackurl,omitempty"`   // fallback_url
	FallbackLogoUrl  string   `json:"fallbacklogo,omitempty"`  // fallback_logo_url
	FallbackTitle    string   `json:"fallbacktitle,omitempty"` // fallback_title
}
