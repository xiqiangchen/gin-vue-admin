package dsp

type ApiGroup struct {
	BidApi
	ImpressionApi
	ClickApi
}

var ApiGroupApp = new(ApiGroup)
