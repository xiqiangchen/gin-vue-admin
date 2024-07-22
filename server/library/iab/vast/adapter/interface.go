package adapter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/library/iab/vast/adapter/base"
)

type VASTHelper interface {
	GetTrackInfo(adm string) (vi *base.VASTInfo, err error)
	SetTrackInfo(adm string, vi *base.VASTInfo) (result string, err error)
	GetProtcol() []int
}
