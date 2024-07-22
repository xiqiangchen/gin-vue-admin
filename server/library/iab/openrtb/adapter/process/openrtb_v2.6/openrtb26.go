package openrtb_v2_5

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/library/iab/openrtb/adapter/base"
	openrtb_v2_62 "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/openrtb2/openrtb_v2.6"
)

type OpenRTB26 struct {
	base.Common
}

func (o *OpenRTB26) SerializationType() int {
	return base.OPEN_RTB_SERIALIZATION_JSON
}
func (o *OpenRTB26) ProtocalVersion() string {
	return base.OPEN_RTB_VERSION_2_6
}

func (o *OpenRTB26) TransformReqTo(ori interface{}) (req *openrtb_v2_62.BidRequest) {
	return ori.(*openrtb_v2_62.BidRequest)
}
func (o *OpenRTB26) TransformReqFrom(ori *openrtb_v2_62.BidRequest) (req interface{}) {
	return ori
}
func (o *OpenRTB26) TransformRespTo(ori interface{}) (resp *openrtb_v2_62.BidResponse) {
	return ori.(*openrtb_v2_62.BidResponse)
}
func (o *OpenRTB26) TransformRespFrom(ori *openrtb_v2_62.BidResponse) (resp interface{}) {
	return ori
}

func (o *OpenRTB26) UnmarshalRequest(data []byte) (result interface{}, err error) {
	req := new(openrtb_v2_62.BidRequest)
	err = json.Unmarshal(data, req)

	// 对方可能通过ext来传的supplychain
	if source := req.Source; source != nil && source.SupplyChain == nil {
		var ext map[string]json.RawMessage
		json.Unmarshal(source.Ext, &ext)
		if c, ok := ext["schain"]; ok {
			source.SupplyChain = new(openrtb_v2_62.SupplyChain)
			json.Unmarshal(c, source.SupplyChain)
		}
	}
	return req, err
}
func (o *OpenRTB26) MarshalRequest(req interface{}) ([]byte, error) {
	return json.Marshal(req)
}
func (o *OpenRTB26) UnmarshalResponse(data []byte) (result interface{}, err error) {
	resp := new(openrtb_v2_62.BidResponse)
	err = json.Unmarshal(data, resp)
	return resp, err
}
func (o *OpenRTB26) MarshalResponse(resp interface{}) ([]byte, error) {
	return json.Marshal(resp)
}

func (o *OpenRTB26) NormalizeRequest(data []byte) (*openrtb_v2_62.BidRequest, error) {
	i, err := o.UnmarshalRequest(data)
	if err != nil {
		return nil, err
	}
	return i.(*openrtb_v2_62.BidRequest), nil

}
func (o *OpenRTB26) DenormalizeRequest(req *openrtb_v2_62.BidRequest) ([]byte, error) {
	return o.MarshalRequest(req)
}
func (o *OpenRTB26) NormalizeResponse(data []byte) (*openrtb_v2_62.BidResponse, error) {
	i, err := o.UnmarshalResponse(data)
	if err != nil {
		return nil, err
	}
	return i.(*openrtb_v2_62.BidResponse), nil
}
func (o *OpenRTB26) DenormalizeResponse(resp *openrtb_v2_62.BidResponse) ([]byte, error) {
	return o.MarshalResponse(resp)
}
