package dsp

type Metrics struct {
	// 指标
	Bid        int `json:"bid,omitempty"`        // 竞价请求
	Offer      int `json:"offer,omitempty"`      // 出价
	Win        int `json:"win,omitempty"`        // 竞得
	Impression int `json:"impression,omitempty"` // 曝光
	Click      int `json:"click,omitempty"`      // 点击
	Landing    int `json:"landing,omitempty"`    // 落地页

	WakeUp   int `json:"wake_up,omitempty"`  // 唤醒
	Active   int `json:"active,omitempty"`   // 激活
	Register int `json:"register,omitempty"` // 注册
	Purchase int `json:"purchase,omitempty"` // 购买

	// 价格相关
	PriceSrc string  `form:"pr,omitempty"` // 结算结果
	Price    float64 `json:"price,omitempty"`
	BidFloor float64 `json:"bid_floor,omitempty" form:"bf,omitempty"` // 底价
}
