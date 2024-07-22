package adapter

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/library/iab/vast/adapter/base"
	"github.com/flipped-aurora/gin-vue-admin/server/library/iab/vast/adapter/process/vast_v2"
	"github.com/flipped-aurora/gin-vue-admin/server/library/iab/vast/adapter/process/vast_v4"
)

func init() {
	InitHelpers()
}

// 使用数组方式代替map提高效率
var Helpers = make([]VASTHelper, 0, 16)

// in:protocal+[]byte   out:VASTInfo
//func

func GetTrackInfo(adm string, protocol int) (vi *base.VASTInfo, err error) {
	// todo 临时只按照2.0 解析
	if protocol >= base.VAST20Wrapper && protocol%2 == 0 {
		protocol = base.VAST20Wrapper
	} else {
		protocol = base.VAST20
	}

	for _, a := range Helpers {
		for _, p := range a.GetProtcol() {
			if p == protocol {
				return a.GetTrackInfo(adm)
			}
		}
	}
	return nil, fmt.Errorf("not implement")
}

func SetTrackInfo(adm string, vi *base.VASTInfo, protocol int) (result string, err error) {
	for _, a := range Helpers {
		for _, p := range a.GetProtcol() {
			if p == protocol {
				return a.SetTrackInfo(adm, vi)
			}
		}
	}
	return "", fmt.Errorf("not implement")
}

func InitHelpers() error {
	Helpers = []VASTHelper{
		&vast_v2.VAST2Helper{},
		&vast_v4.VAST4Helper{},
	}
	return nil
}
