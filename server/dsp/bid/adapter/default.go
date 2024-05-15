package adapter

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/constant"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp/bid"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

var _ Adapter = (*defaultAdapter)(nil)

type defaultAdapter struct {
	AdxId int
}

func NewDefaultAdapter(adxId ...int) *defaultAdapter {
	if len(adxId) > 0 {
		return &defaultAdapter{AdxId: adxId[0]}
	}
	return &defaultAdapter{AdxId: constant.DefaultAdxId}

}

func (d *defaultAdapter) From(header http.Header, byt []byte) (*bid.BidRequest, error) {
	req := new(bid.BidRequest)
	// 暂不处理压缩
	contentType := strings.ToLower(header.Get("content-type"))
	switch {
	case strings.Contains(contentType, "application/json"):
		if err := json.Unmarshal(byt, req); err != nil {
			global.GVA_LOG.Error("请求解释失败", zap.Error(err))
			return nil, err
		}
	default:
		if err := proto.Unmarshal(byt, req); err != nil {
			global.GVA_LOG.Error("请求解释失败", zap.Error(err))
			return nil, err
		}
	}

	return req, nil
}

func (d *defaultAdapter) To(response *bid.BidResponse) (any, error) {
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
