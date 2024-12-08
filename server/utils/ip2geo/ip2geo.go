package ip2geo

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/ip2geo/trans"
	"strconv"
	"strings"
)

var TranIpResultEmpty = errors.New("ip2geo Parse is return null")

// LoadIp2Geo 加载IP库离线文件
func LoadIp2Geo(ipv4File, ipv6File string) error {
	if len(ipv4File) > 0 {
		if err := ipv4Loader.Load(ipv4File); err != nil {
			return err
		}
	}
	if len(ipv6File) > 0 {
		if err := ipv6Loader.Load(ipv6File); err != nil {
			return err
		}
	}
	return nil
}

// Parse IP转成location表的省市区ID
func Parse(ip string) (*trans.IpTrans, error) {
	if len(ip) == 0 {
		return nil, fmt.Errorf("ip2geo Parse ip is null")
	}
	tran, err := GetIp2Geo(ip)
	if err != nil {
		return nil, fmt.Errorf("ip2geo Parse GetIp2Geo ip:%s;err:%s", ip, err.Error())
	}

	if len(tran.City) == 0 && len(tran.Country) == 0 && len(tran.Province) == 0 {
		return nil, TranIpResultEmpty
	}

	return tran, nil
}

// GetIp2Geo IP转离线库的地址信息
func GetIp2Geo(ip string) (*trans.IpTrans, error) {
	var (
		ipTrans *trans.IpTrans
		ipStr   string
		err     error
	)
	if strings.Contains(ip, ":") {
		ipStr, err = ipv6Loader.Get(ip)
	} else {
		ipStr, err = ipv4Loader.Get(ip)
	}
	if err != nil {
		return nil, fmt.Errorf("ip2geo err %s,ip:%s", err.Error(), ip)
	}
	ipSplit := strings.Split(ipStr, "|")
	if len(ipSplit) < 15 {
		return nil, fmt.Errorf("ip2geo ip:%s,ip result is less than 14 param", ip)
	}

	// 字段顺序以model文件为准
	ipTrans = &trans.IpTrans{
		Continent:     ipSplit[0],
		Country:       ipSplit[1],
		CountryEn:     ipSplit[2],
		Province:      ipSplit[3],
		ProvinceEn:    ipSplit[4],
		City:          ipSplit[5],
		CityEn:        ipSplit[6],
		CountryCode:   ipSplit[9],
		ZipCode:       ipSplit[10],
		Carrier:       ipSplit[11],
		LbsCityId:     ipSplit[12],
		TimeZone:      ipSplit[13],
		LbsProvinceId: ipSplit[14],
	}

	longitude, err := strconv.ParseFloat(ipSplit[7], 64)
	if err == nil {
		ipTrans.Longitude = longitude
	}

	latitude, err := strconv.ParseFloat(ipSplit[8], 64)
	if err == nil {
		ipTrans.Latitude = latitude
	}

	return ipTrans, nil
}
