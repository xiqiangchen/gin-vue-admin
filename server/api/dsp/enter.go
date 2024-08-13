package dsp

type ApiGroup struct {
	BidApi
	ImpressionApi
	ClickApi
	LandingApi
}

var ApiGroupApp = new(ApiGroup)
