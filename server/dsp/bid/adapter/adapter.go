package adapter

import "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/bid"

type Adapter interface {
	From(byt []byte) (*bid.BidRequest, error)
	To(response *bid.BidResponse) ([]byte, error)
	GetAdxId() int
	GetProtocol() int
	GetAdxPriceMacro() string
}

func GetAdapter(adxId int) Adapter {
	return NewDefaultAdapter()
}
