package ip2geo

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/ip2geo/trans"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/ip2geo/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var ipv4DataPath string
var ipv6DataPath string
var ipv4EnDataPath string
var ipv6EnDataPath string

func init() {
	ipv4DataPath = strings.Replace(getWorkPath(), utils.WindowsDirSpitReplace("/utils/ip2geo"), "", 1) + "/resource/data/ip/ipdatacloud_city_v4_cn.dat"
	ipv4DataPath = utils.WindowsDirSpitReplace(ipv4DataPath)

	ipv6DataPath = strings.Replace(getWorkPath(), utils.WindowsDirSpitReplace("/utils/ip2geo"), "", 1) + "/resource/data/ip/ipdatacloud_city_v6_cn.dat"
	ipv6DataPath = utils.WindowsDirSpitReplace(ipv6DataPath)

	ipv4EnDataPath = strings.Replace(getWorkPath(), utils.WindowsDirSpitReplace("/utils/ip2geo"), "", 1) + "/resource/data/ip/ipdatacloud_city_v4_en.dat"
	ipv4EnDataPath = utils.WindowsDirSpitReplace(ipv4EnDataPath)

	ipv6EnDataPath = strings.Replace(getWorkPath(), utils.WindowsDirSpitReplace("/utils/ip2geo"), "", 1) + "/resource/data/ip/ipdatacloud_city_v6_en.dat"
	ipv6EnDataPath = utils.WindowsDirSpitReplace(ipv6EnDataPath)

}

func TestIpv4CN(t *testing.T) {

	if err := ipv4Loader.Load(ipv4DataPath); err != nil {
		t.Fatalf("can not load data file")
	}

	str, _ := ipv4Loader.Get("220.162.31.109")
	ipSplit := strings.Split(str, "|")
	assert.NotEqual(t, ipSplit, nil)
	if len(ipSplit) < 14 {
		t.Logf("ip:%s, is less than 14", ipSplit)
	}
	ips := &trans.IpTrans{
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

	t.Logf("%+v \n", ips)
}

func TestIpv6CN(t *testing.T) {
	ipv6s := []string{
		"2409:890e:a00:1075:b1a0:58e:2c53:9d9c",
		//"2409:8a28:6a96:8a60:d902:8da4:b781:8138",
		"2409:8a30:5a0:6301:159e:a65c:1f04:9357",
		"2409:8934:1097:8281:f09f:7991:5f5e:d062",
		"240e:361:2e2e:be00:89b0:7c82:65cf:18d5",
		"2408:8256:3486:b9c:8975:fea6:4bb2:ee6e",
		"2408:8414:2640:1b5f:1:1:da6b:450a",
		"2409:891f:9204:58c:c5c9:c971:896c:cd9d",
		"2409:8907:833a:e59:e8b9:9e09:6da0:f23f",
		"2409:896a:9dfa:a55:c0ca:15ff:fea7:6e56",
	}

	if err := ipv6Loader.Load(ipv6DataPath); err != nil {
		t.Fatalf("can not load data file")
	}

	for _, ipv6 := range ipv6s {
		str, _ := ipv6Loader.Get(ipv6)
		ipSplit := strings.Split(str, "|")
		assert.NotEqual(t, ipSplit, nil)
		if len(ipSplit) < 14 {
			t.Logf("ip:%s, is less than 14", ipSplit)
			continue
		}
		ipInfo := &trans.IpTrans{
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

		fmt.Printf("ip:%s, info %#v", ipv6, ipInfo)
		fmt.Println()
	}

}

func TestIpv4EN(t *testing.T) {

	if err := ipv4Loader.Load(ipv4EnDataPath); err != nil {
		t.Fatalf("can not load data file")
	}

	str, _ := ipv4Loader.Get("34.138.58.111")
	ipSplit := strings.Split(str, "|")
	assert.NotEqual(t, ipSplit, nil)
	if len(ipSplit) < 14 {
		t.Logf("ip:%s, is less than 14", ipSplit)
	}
	ips := &trans.IpTrans{
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

	t.Logf("%+v \n", ips)
}

func TestIpv6EN(t *testing.T) {
	ipv6s := []string{
		"2405:204:911c:60ff:ffff:ffff:ffff:ffff",
		"2a01:e0a:556:bfff:ffff:ffff:ffff:ffff",
		"2a13:85c7:ffff:ffff:ffff:ffff:ffff:ffff",
		"2a13:793f:ffff:ffff:ffff:ffff:ffff:ffff",
		//"2a13:7947:ffff:ffff:ffff:ffff:ffff:ffff",
		"2001:718:1e04:ffff:ffff:ffff:ffff:ffff",
		"2001:718:1fff:ffff:ffff:ffff:ffff:ffff",
		"2a02:8108:9bbf:ffff:ffff:ffff:ffff:ffff",
		"2a02:8108:9c3e:ffff:ffff:ffff:ffff:ffff",
		"2a11:bbc7:ffff:ffff:ffff:ffff:ffff:ffff",
	}

	if err := ipv6Loader.Load(ipv6EnDataPath); err != nil {
		t.Fatalf("can not load data file")
	}

	for _, ipv6 := range ipv6s {
		str, _ := ipv6Loader.Get(ipv6)
		ipSplit := strings.Split(str, "|")
		assert.NotEqual(t, ipSplit, nil)
		if len(ipSplit) < 14 {
			t.Logf("ip:%s, is less than 14", ipSplit)
			continue
		}
		ipInfo := &trans.IpTrans{
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

		fmt.Printf("ip:%s, info %#v", ipv6, ipInfo)
		fmt.Println()
	}

}

func getWorkPath() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}
