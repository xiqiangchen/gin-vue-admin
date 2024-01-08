package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CreativeSearch struct {
	ad.Creative
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type CreativeBatch struct {
	ad.Creative
	Images []*Material `json:"images" form:"images"`
	Videos []*Material `json:"videos" form:"videos"`
}

type Material struct {
	Id  int    `json:"id" form:"id"`
	Url string `json:"url" form:"url"`
}
