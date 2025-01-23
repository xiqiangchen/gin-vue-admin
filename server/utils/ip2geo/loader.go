package ip2geo

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/ip2geo/trans"
)

type IIp2Geo interface {
	Get(ip string) (string, error)
	Load(fileName string) error
}

var ipv4Loader IIp2Geo = &IPv4Loader{}
var ipv6Loader IIp2Geo = &IPv6Loader{}

type IPv4Loader struct {
	ipInfo *trans.IpInfo
}

func (loader *IPv4Loader) Load(fileName string) (err error) {
	loader.ipInfo, err = trans.GetObject(fileName)
	if err != nil {
		return fmt.Errorf("error to load ipv4 dat file %s", err.Error())
	}
	return
}

func (loader *IPv4Loader) Get(ip string) (string, error) {
	return loader.ipInfo.Get(ip)
}

type IPv6Loader struct {
	ipInfo *trans.IpInfoV6
}

func (loader *IPv6Loader) Load(fileName string) (err error) {
	loader.ipInfo, err = trans.GetDistrictObjectV6(fileName)
	if err != nil {
		return fmt.Errorf("error to load ipv4 datipv6 file %s", err.Error())
	}
	return
}

func (loader *IPv6Loader) Get(ip string) (string, error) {
	return loader.ipInfo.Get(ip)
}
