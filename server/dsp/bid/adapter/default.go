package adapter

import "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/bid"

var _ Adapter = (*defaultAdapter)(nil)

type defaultAdapter struct {
}

func NewDefaultAdapter() *defaultAdapter {
	return &defaultAdapter{}
}

func (d *defaultAdapter) From(byt []byte) (bid.BidRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (d *defaultAdapter) To(response bid.BidResponse) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (d *defaultAdapter) GetAdxId() int {
	//TODO implement me
	panic("implement me")
}

func (d *defaultAdapter) GetProtocol() int {
	//TODO implement me
	panic("implement me")
}

func (d *defaultAdapter) GetAdxPriceMacro() string {
	//TODO implement me
	panic("implement me")
}
