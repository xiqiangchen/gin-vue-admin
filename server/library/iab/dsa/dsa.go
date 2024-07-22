package dsa

// DigitalServicesAct
type DSARequest struct {
	DSARequired  int            `json:"dsarequired"`
	Pubrender    int            `json:"pubrender"`
	Datatopub    int            `json:"datatopub"`
	Transparency []Transparency `json:"transparency"`
}

type DSAResponse struct {
	Behalf       string         `json:"behalf"`       // 展示广告的主体名称（要放舜飞还是dsp方的广告主？）
	Paid         string         `json:"paid"`         // 广告支付者 （一般和上面的 behalf 是相同的）
	Transparency []Transparency `json:"transparency"` //
	Adrender     int            `json:"adrender"`     //
}

type Transparency struct {
	Domain    string `json:"domain"`
	Dsaparams []int  `json:"dsaparams"`
}
