package adapter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp/bid"
	"net/http"
)

type Adapter interface {
	From(http.Header, []byte) (*bid.BidRequest, error)
	To(response *bid.BidResponse) (any, error)
	GetAdxId() int
	GetProtocol() int
	GetAdxPriceMacro() string
}

func GetAdapter(adxId int) Adapter {
	return NewDefaultAdapter()
}
