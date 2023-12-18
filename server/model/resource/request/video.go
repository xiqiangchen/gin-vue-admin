package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/resource"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type VideoSearch struct{
    resource.Video
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
