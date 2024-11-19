package adapter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/dsp/bid/pricer"
	protocol "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/openrtb2/openrtb_v2.6"
	"net/http"
)

type Adapter interface {
	From(http.Header, []byte) (*protocol.BidRequest, error)
	To(response *protocol.BidResponse) (any, error)
	GetAdxId() int
	GetProtocol() int
	GetAdxPriceMacro() string
	GetPricer() *pricer.Decrypter
}

func GetAdapter(adxId int) Adapter {
	return NewDefaultAdapter()
}
