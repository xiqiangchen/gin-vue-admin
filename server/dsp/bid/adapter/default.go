package adapter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/dsp"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp/bid"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
)

var _ Adapter = (*defaultAdapter)(nil)

type defaultAdapter struct {
	AdxId int
}

func NewDefaultAdapter(adxId ...int) *defaultAdapter {
	if len(adxId) > 0 {
		return &defaultAdapter{AdxId: adxId[0]}
	}
	return &defaultAdapter{AdxId: dsp.DefaultAdxId}

}

func (d *defaultAdapter) From(byt []byte) (*bid.BidRequest, error) {
	req := new(bid.BidRequest)
	if err := proto.Unmarshal(byt, req); err != nil {
		global.GVA_LOG.Error("请求解释失败", zap.Error(err))
		return nil, err
	}

	return req, nil
}

func (d *defaultAdapter) To(response *bid.BidResponse) ([]byte, error) {
	return response.Marshal()
}

func (d *defaultAdapter) GetAdxId() int {
	return d.AdxId
}

func (d *defaultAdapter) GetProtocol() int {
	return dsp.DefaultProtocol
}

func (d *defaultAdapter) GetAdxPriceMacro() string {
	return dsp.DefaultPriceMacro
}
