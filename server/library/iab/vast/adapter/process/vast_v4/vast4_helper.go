package vast_v4

import "github.com/flipped-aurora/gin-vue-admin/server/library/iab/vast/adapter/base"

type VAST4Helper struct {
	base.Common
}

func (h *VAST4Helper) GetTrackInfo(adm string) (vi *base.VASTInfo, err error) {
	return
}
func (h *VAST4Helper) SetTrackInfo(adm string, vi *base.VASTInfo) (result string, err error) {
	return
}

func (h *VAST4Helper) GetProtcol() []int {
	return []int{base.VAST40, base.VAST40Wrapper}
}
