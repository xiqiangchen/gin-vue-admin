package adapter

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/native1/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/native1/response"
	"strconv"
	"strings"
)

// todo domark 是否直接使用string 删掉/加上 "{"native":   }" 就好
//func init() {
//	pattern := `{.*("native".*):`
//	re := regexp.MustCompile(pattern)
//}

func UnmarshalRequest(adMarkup string, version string) (req *request.Request, err error) {
	if strings.HasPrefix(adMarkup, `"`) {
		adMarkup, _ = strconv.Unquote(adMarkup)
	}
	native12 := new(request.Request)
	_adMarkup := adMarkup
	if len(_adMarkup) > 30 {
		_adMarkup = _adMarkup[:30]
	}
	if strings.Contains(_adMarkup, "native") {
		native10 := new(OuterLayerExt)
		native10.Native = native12
		err = json.Unmarshal([]byte(adMarkup), native10)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(adMarkup), native12)
		if err != nil {
			return nil, err
		}
	}
	return native12, nil
}

func MarshalRequest(req *request.Request, version string) (adMarkup string, err error) {
	if version == NATIVE_1_0 {
		native10 := new(OuterLayerExt)
		native10.Native = req
		msg, err := json.Marshal(native10)
		if err != nil {
			return "", err
		}
		return string(msg), nil
	} else {
		msg, err := json.Marshal(req)
		if err != nil {
			return "", err
		}
		return string(msg), nil
	}
}

func MarshalResponse(resp *response.Response) (adMarkup string, err error) {
	//if version == NATIVE_1_0 {
	//	native10 := new(OuterLayerExt)
	//	native10.Native = resp
	//	msg, err := json.Marshal(native10)
	//	if err != nil {
	//		return "", err
	//	}
	//	return string(msg), nil
	//} else {
	//	resp = new(response.Response)
	//	msg, err := json.Marshal(resp)
	//	if err != nil {
	//		return "", err
	//	}
	//	return string(msg), err
	//}
	// 1.2 或 1.1 或 1.0 都是 "{"native":xxxx}",不需要指定版本
	ole := new(OuterLayerExt)
	ole.Native = resp
	msg, err := json.Marshal(ole)
	if err != nil {
		return "", err
	}
	return string(msg), err
}
func UnmarshalResponse(adMarkup string) (resp *response.Response, err error) {
	if strings.HasPrefix(adMarkup, `"`) {
		adMarkup, _ = strconv.Unquote(adMarkup)
	}
	resp = new(response.Response)
	if strings.Contains(adMarkup[:30], "native") {
		ole := new(OuterLayerExt)
		ole.Native = resp
		err = json.Unmarshal([]byte(adMarkup), ole)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(adMarkup), resp)
		if err != nil {
			return nil, err
		}
	}
	return resp, err
}

const (
	NATIVE_1_0 = "1.0"
	NATIVE_1_2 = "1.2"
)

// 临时加上个1.0的外层
type OuterLayerExt struct {
	Native interface{} `json:"native,omitempty"`
}
