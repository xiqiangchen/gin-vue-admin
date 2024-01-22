package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/assert"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type TargetSearch struct {
	assert.Target
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
