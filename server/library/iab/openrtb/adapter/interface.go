package adapter

import (
	openrtb_v2_62 "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/openrtb2/openrtb_v2.6"
)

type OpenRtbAdapter interface {
	SerializationType() int
	ProtocalVersion() string

	TransformReqTo(interface{}) *openrtb_v2_62.BidRequest
	TransformReqFrom(*openrtb_v2_62.BidRequest) interface{}
	TransformRespTo(interface{}) *openrtb_v2_62.BidResponse
	TransformRespFrom(response *openrtb_v2_62.BidResponse) interface{}

	// 将响应序列化 对json 没什么用，主要是用于proto类
	UnmarshalRequest(msg []byte) (interface{}, error)
	MarshalRequest(interface{}) ([]byte, error)
	UnmarshalResponse([]byte) (interface{}, error)
	MarshalResponse(interface{}) ([]byte, error)

	NormalizeRequest([]byte) (*openrtb_v2_62.BidRequest, error)
	DenormalizeRequest(*openrtb_v2_62.BidRequest) ([]byte, error)
	NormalizeResponse([]byte) (*openrtb_v2_62.BidResponse, error)
	DenormalizeResponse(*openrtb_v2_62.BidResponse) ([]byte, error)
}

type Ext struct {
}
