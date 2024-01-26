package bid

import "github.com/flipped-aurora/gin-vue-admin/server/model/ad"

type Plan struct {
	ad.Plan
	RealCampaigns      []*ad.Campaign `json:"real_campaigns"`
	VirtuallyCampaigns []*ad.Campaign `json:"virtually_campaigns"`
}
