package adapter

import (
	"fmt"
	adp_v2_5 "github.com/flipped-aurora/gin-vue-admin/server/library/iab/openrtb/adapter/process/openrtb_v2.5"
	adp_v2_6 "github.com/flipped-aurora/gin-vue-admin/server/library/iab/openrtb/adapter/process/openrtb_v2.6"
	openrtb_v2_62 "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/openrtb2/openrtb_v2.6"
)

// todo 临时，后续把初始化放到项目启动流程里
func init() {
	InitAdapters()
}

// 使用数组方式代替map提高效率
var Adapters = make([]OpenRtbAdapter, 0, 16)

// 1.传入版本的 []byte 转换成内部 ortb （req 和 resp 版本）
// 2.传入内部ortb 转换成 []byte (req 和 resp 版本)
func RequestAdapter(version string, serialization int, body []byte) (req *openrtb_v2_62.BidRequest, err error) {
	for _, a := range Adapters {
		if a.ProtocalVersion() == version && a.SerializationType() == serialization {
			return a.NormalizeRequest(body)
		}
	}
	return nil, fmt.Errorf("not implement")
}
func RequestAdapter2(version string, serialization int, req *openrtb_v2_62.BidRequest) (data []byte, err error) {
	for _, a := range Adapters {
		if a.ProtocalVersion() == version && a.SerializationType() == serialization {
			return a.DenormalizeRequest(req)
		}
	}
	return nil, fmt.Errorf("not implement")
}
func ResponseAdapter(version string, serialization int, body []byte) (resp *openrtb_v2_62.BidResponse, err error) {
	for _, a := range Adapters {
		if a.ProtocalVersion() == version && a.SerializationType() == serialization {
			return a.NormalizeResponse(body)
		}
	}
	return nil, fmt.Errorf("not implement")
}
func ResponseAdapter2(version string, serialization int, resp *openrtb_v2_62.BidResponse) (data []byte, err error) {
	for _, a := range Adapters {
		if a.ProtocalVersion() == version && a.SerializationType() == serialization {
			return a.DenormalizeResponse(resp)
		}
	}
	return nil, fmt.Errorf("not implement")
}

func GetAdapter(version string, serialization int) OpenRtbAdapter {
	for _, a := range Adapters {
		if a.ProtocalVersion() == version && a.SerializationType() == serialization {
			return a
		}
	}
	return nil
}

func InitAdapters() error {
	Adapters = []OpenRtbAdapter{
		&adp_v2_5.OpenRTB25{},
		&adp_v2_6.OpenRTB26{},
	}
	return nil
}
