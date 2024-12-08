package tester

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/ip2geo/csv"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/ip2geo/trans"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/ip2geo/utils"
	"strings"
	"testing"
)

var (
	ipv4DataPath   string
	ipv6DataPath   string
	ipv4EnDataPath string
	ipv6EnDataPath string
)

func init() {
	ipv4DataPath = utils.InitDataPath(utils.WindowsDirSpitReplace("library/utils/ip2geo/tester"), "deploy/tx/data/resources/ip2geo/ipdatacloud_city_v4_cn.dat")
	ipv6DataPath = utils.InitDataPath(utils.WindowsDirSpitReplace("library/utils/ip2geo/tester"), "deploy/tx/data/resources/ip2geo/ipdatacloud_city_v6_cn.dat")
	ipv4EnDataPath = utils.InitDataPath(utils.WindowsDirSpitReplace("library/utils/ip2geo/tester"), "deploy/tx/data/resources/ip2geo/ipdatacloud_city_v4_en.dat")
	ipv6EnDataPath = utils.InitDataPath(utils.WindowsDirSpitReplace("library/utils/ip2geo/tester"), "deploy/tx/data/resources/ip2geo/ipdatacloud_city_v6_en.dat")
}

func TestIpv4(t *testing.T) {
	dataObj, _ := trans.GetObject(ipv4DataPath)
	ipv4List, err := csv.GetAdxIpList(csv.Ipv4, 1000)
	if err != nil {
		t.Errorf(err.Error())
	}
	for i, ipData := range ipv4List {
		str, err := dataObj.Get(ipData.IP)
		if err != nil {
			t.Errorf("index %d err %s,ip:%s", i, err.Error(), ipData.IP)
			continue
		}
		ipSplit := strings.Split(str, "|")
		if len(ipSplit) < 14 {
			t.Errorf("index %d err %s,ip:%s", i, fmt.Errorf("ip result is less than 14 param"), ipData.IP)
			continue
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
		t.Logf("ip:%s;%+v \n", ipData.IP, ips)
	}
}

func TestIpv6(t *testing.T) {
	dataObj, _ := trans.GetDistrictObjectV6(ipv6DataPath)
	ipv6List, err := csv.GetAdxIpList(csv.Ipv6, 3)
	if err != nil {
		t.Errorf(err.Error())
	}
	for i, ipData := range ipv6List {
		str, err := dataObj.Get(ipData.IP)
		if err != nil {
			t.Errorf("index %d err %s,ip:%s", i, err.Error(), ipData.IP)
			continue
		}
		ipSplit := strings.Split(str, "|")
		if len(ipSplit) < 14 {
			t.Errorf("index %d err %s,ip:%s", i, fmt.Errorf("ip result is less than 14 param"), ipData.IP)
			continue
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
		t.Logf("ip:%s;%+v \n", ipData.IP, ips)
	}
}

// 1条 ipv4的benchmark
func BenchmarkIpv4SingleIp(b *testing.B) {
	dataObj, _ := trans.GetObject(ipv4DataPath)
	ipv4List, err := csv.GetAdxIpList(csv.Ipv4, 1)
	if err != nil {
		b.Errorf(err.Error())
	}

	// 设置并行的goroutine数量
	b.SetParallelism(2)

	// 开始计算内存分配
	b.ReportAllocs()

	// 重置计时器
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for _, ipData := range ipv4List {
			_, err := dataObj.Get(ipData.IP)
			if err != nil {
				if strings.Contains(err.Error(), "不合法") {
					continue
				}
				b.Errorf("index %d err %s,ip:%s", n, err.Error(), ipData.IP)
				continue
			}
		}
	}

	//goos: windows
	//goarch: amd64
	//pkg: github.com/flipped-aurora/gin-vue-admin/server/utils/ip2geo/tester
	//BenchmarkIpv4SingleIp-8          1584812               769 ns/op             208
	// B/op          3 allocs/op
	//PASS

	//1584812： 这是在基准测试期间运行的迭代次数。这个数量是由Go测试框架自动决定的，以确保基准测试的结果是可靠的。
	//769 ns/op： 这表示每次迭代大约需要769纳秒。这是你的代码的性能度量，表示每次操作所需的时间。
	//208 B/op： 这表示每次迭代大约需要分配208字节的内存。这是你的代码的内存使用度量，表示每次操作所需的内存。
	//3 allocs/op： 这表示每次迭代大约需要进行3次内存分配。这是你的代码的内存分配次数度量，表示每次操作所分配的内存块数量。
}

// 1条 ipv6的benchmark
func BenchmarkIpv6SingleIp(b *testing.B) {
	dataObj, _ := trans.GetDistrictObjectV6(ipv6DataPath)
	ipList, err := csv.GetAdxIpList(csv.Ipv6, 1)
	if err != nil {
		b.Errorf(err.Error())
	}

	// 设置并行的goroutine数量
	b.SetParallelism(2)

	// 开始计算内存分配
	b.ReportAllocs()

	// 重置计时器
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for _, ipData := range ipList {
			_, err := dataObj.Get(ipData.IP)
			if err != nil {
				if strings.Contains(err.Error(), "不合法") {
					continue
				}
				b.Errorf("index %d err %s,ip:%s", n, err.Error(), ipData.IP)
				continue
			}
		}
	}

	//goos: windows
	//goarch: amd64
	//pkg: github.com/flipped-aurora/gin-vue-admin/server/utils/ip2geo/tester
	//BenchmarkIpv6SingleIp-8            44448             25180 ns/op            4056
	// B/op         95 allocs/op
	//PASS

	//44448： 这是在基准测试期间运行的迭代次数。这个数量是由Go测试框架自动决定的，以确保基准测试的结果是可靠的。
	//25180 ns/op： 这表示每次迭代大约需要25180纳秒，即0.02518 毫秒。这是你的代码的性能度量，表示每次操作所需的时间。
	//4056 B/op： 这表示每次迭代大约需要分配4056字节的内存。这是你的代码的内存使用度量，表示每次操作所需的内存。
	//95 allocs/op： 这表示每次迭代大约需要进行95次内存分配。这是你的代码的内存分配次数度量，表示每次操作所分配的内存块数量。
}

// 1条 ipv4的benchmark
func BenchmarkIpv4EnSingleIp(b *testing.B) {
	dataObj, _ := trans.GetObject(ipv4EnDataPath)
	ipv4List := []string{
		"204.13.158.178",
	}

	// 设置并行的goroutine数量
	b.SetParallelism(2)

	// 开始计算内存分配
	b.ReportAllocs()

	// 重置计时器
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for _, ip := range ipv4List {
			_, err := dataObj.Get(ip)
			if err != nil {
				if strings.Contains(err.Error(), "不合法") {
					continue
				}
				b.Errorf("index %d err %s,ip:%s", n, err.Error(), ip)
				continue
			}
		}
	}

	//goos: windows
	//goarch: amd64
	//pkg: github.com/flipped-aurora/gin-vue-admin/server/utils/ip2geo/tester
	//BenchmarkIpv4EnSingleIp-8        1972741               573 ns/op             224
	// B/op          3 allocs/op
	//PASS

	//1972741： 这是在基准测试期间运行的迭代次数。这个数量是由Go测试框架自动决定的，以确保基准测试的结果是可靠的。
	//573 ns/op： 这表示每次迭代大约需要573纳秒。这是你的代码的性能度量，表示每次操作所需的时间。
	//224 B/op： 这表示每次迭代大约需要分配224字节的内存。这是你的代码的内存使用度量，表示每次操作所需的内存。
	//3 allocs/op： 这表示每次迭代大约需要进行3次内存分配。这是你的代码的内存分配次数度量，表示每次操作所分配的内存块数量。
}

// 1条 ipv6的benchmark
func BenchmarkIpv6EnSingleIp(b *testing.B) {
	dataObj, _ := trans.GetDistrictObjectV6(ipv6EnDataPath)
	ipList := []string{
		"2405:204:911c:60ff:ffff:ffff:ffff:ffff",
	}

	// 设置并行的goroutine数量
	b.SetParallelism(2)

	// 开始计算内存分配
	b.ReportAllocs()

	// 重置计时器
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for _, ip := range ipList {
			_, err := dataObj.Get(ip)
			if err != nil {
				if strings.Contains(err.Error(), "不合法") {
					continue
				}
				b.Errorf("index %d err %s,ip:%s", n, err.Error(), ip)
				continue
			}
		}
	}

	//goos: windows
	//goarch: amd64
	//pkg: github.com/flipped-aurora/gin-vue-admin/server/utils/ip2geo/tester
	//BenchmarkIpv6EnSingleIp-8          39192             29513 ns/op            3840
	// B/op         89 allocs/op
	//PASS

	//39192： 这是在基准测试期间运行的迭代次数。这个数量是由Go测试框架自动决定的，以确保基准测试的结果是可靠的。
	//29513 ns/op： 这表示每次迭代大约需要29513纳秒，即0.0295 毫秒。这是你的代码的性能度量，表示每次操作所需的时间。
	//3840 B/op： 这表示每次迭代大约需要分配3840字节的内存。这是你的代码的内存使用度量，表示每次操作所需的内存。
	//89 allocs/op： 这表示每次迭代大约需要进行89次内存分配。这是你的代码的内存分配次数度量，表示每次操作所分配的内存块数量。
}
