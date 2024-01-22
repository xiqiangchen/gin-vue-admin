package bid

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	"github.com/flipped-aurora/gin-vue-admin/server/service/dsp"
	"time"
)

var ActiveCampaigns []*ad.Campaign

func Load() error {

	plans := loadPlans()
	_ = plans

	var campaigns, finalCampaigns []*ad.Campaign
	db := global.GVA_DB.Model(&ad.Campaign{})
	db.Where("status = ? AND is_virtually = ? AND filter = ? AND start_at > ? AND end_at < ?", dsp.StatusOn, dsp.Real, dsp.Pass,
		time.Now(), time.Now())

	_ = db.Preload("Plan").Preload("Creatives").Find(&campaigns).Error

	for _, c := range campaigns {
		_ = c
		if 1 == 1 {
			finalCampaigns = append(finalCampaigns, c)
		}
	}

	return nil
}

func loadPlans() (plans []*ad.Plan) {
	return
}
