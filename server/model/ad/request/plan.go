package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PlanSearch struct {
	ad.Plan
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
