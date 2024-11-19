package adapter

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/constant"
	"github.com/flipped-aurora/gin-vue-admin/server/dsp/bid/pricer"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	protocol "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/openrtb2/openrtb_v2.6"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

var _ Adapter = (*defaultAdapter)(nil)

type defaultAdapter struct {
	AdxId  int
	Pricer *pricer.Decrypter
}

func (d *defaultAdapter) GetPricer() *pricer.Decrypter {
	return d.Pricer
}

func NewDefaultAdapter(adxId ...int) *defaultAdapter {
	if len(adxId) > 0 {
		return &defaultAdapter{AdxId: adxId[0]}
	}
	if pri, err := pricer.New("dac78876796e5479fc66ca252790cb1c3cb2687a04f205acb54ccbdb8504b775",
		"a6b120682f56d2b7df82087b79bee9cd2435bbfa6f2dbf35638aeafc039350ee"); err == nil {
		return &defaultAdapter{AdxId: constant.DefaultAdxId, Pricer: pri}
	} else {
		return &defaultAdapter{AdxId: constant.DefaultAdxId, Pricer: nil}
	}

}

func (d *defaultAdapter) From(header http.Header, byt []byte) (*protocol.BidRequest, error) {
	req := new(protocol.BidRequest)
	// 暂不处理压缩
	contentType := strings.ToLower(header.Get("content-type"))
	switch {
	case strings.Contains(contentType, "application/json"):
		if err := json.Unmarshal(byt, req); err != nil {
			global.GVA_LOG.Error("请求解释失败", zap.Error(err))
			return nil, err
		}
	default:
		if err := json.Unmarshal(byt, req); err != nil {
			global.GVA_LOG.Error("请求解释失败", zap.Error(err))
			return nil, err
		}
	}

	return req, nil
}

func (d *defaultAdapter) To(response *protocol.BidResponse) (any, error) {
	return response, nil
	//return json.Marshal(response)
	//return response.Marshal()
}

func (d *defaultAdapter) GetAdxId() int {
	return d.AdxId
}

func (d *defaultAdapter) GetProtocol() int {
	return constant.DefaultProtocol
}

func (d *defaultAdapter) GetAdxPriceMacro() string {
	return constant.DefaultPriceMacro
}
