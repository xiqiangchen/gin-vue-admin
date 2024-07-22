package base

import (
	openrtb_v2_62 "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/openrtb2/openrtb_v2.6"
)

type Common struct {
}

func (o *Common) SerializationType() int {
	return OPEN_RTB_SERIALIZATION_JSON
}
func (o *Common) ProtocalVersion() string {
	return OPEN_RTB_VERSION_2_6
}

func (o *Common) TransformReqTo(ori interface{}) (req *openrtb_v2_62.BidRequest) {
	return nil
}
func (o *Common) TransformReqFrom(ori *openrtb_v2_62.BidRequest) (req interface{}) {
	return nil
}
func (o *Common) TransformRespTo(ori interface{}) (resp *openrtb_v2_62.BidResponse) {
	return nil
}
func (o *Common) TransformRespFrom(ori *openrtb_v2_62.BidResponse) (resp interface{}) {
	return nil
}

func (o *Common) UnmarshalRequest([]byte) (interface{}, error) {
	return nil, nil
}
func (o *Common) MarshalRequest(interface{}) ([]byte, error) {
	return nil, nil
}
func (o *Common) UnmarshalResponse([]byte) (interface{}, error) {
	return nil, nil
}
func (o *Common) MarshalResponse(interface{}) ([]byte, error) {
	return nil, nil
}

func (o *Common) NormalizeRequest([]byte) (*openrtb_v2_62.BidRequest, error) {
	return nil, nil
}
func (o *Common) DenormalizeRequest(*openrtb_v2_62.BidRequest) ([]byte, error) {
	return nil, nil
}

func (o *Common) NormalizeResponse([]byte) (*openrtb_v2_62.BidResponse, error) {
	return nil, nil
}
func (o *Common) DenormalizeResponse(*openrtb_v2_62.BidResponse) ([]byte, error) {
	return nil, nil
}
