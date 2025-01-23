package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/ip2geo/utils"
	"io"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	IP   string
	ALat float64
	ALon float64
	Lat  float64
	Lon  float64
	Cnt  int
	Rgn  int
	City int
}

const (
	Ipv4 = iota + 1
	Ipv6
	limit int = 50000 // 默认全拿
)

var dataPath string

func init() {
	dataPath = utils.InitDataPath("tester", "csv/adx_ip_data.csv")
}

func GetAdxIpList(ipType int, lmt int) ([]Record, error) {
	if lmt == 0 {
		lmt = limit
	}
	list := make([]Record, 0, lmt)
	adxIps, err := getAdxIpData()
	if err != nil {
		return nil, err
	}
	var matchCount int
	for _, r := range adxIps {
		if matchCount >= lmt {
			break
		}
		if ipType == Ipv4 && !isIPv4(r.IP) {
			continue
		}
		if ipType == Ipv6 && isIPv4(r.IP) {
			continue
		}
		list = append(list, r)
		matchCount++
	}
	return list, nil
}

func isIPv4(address string) bool {
	return strings.Count(address, ":") < 2
}

func getAdxIpData() ([]Record, error) {
	var err error
	list := make([]Record, 0, 50000)
	// 打开文件
	file, err := os.Open(dataPath)
	if err != nil {
		fmt.Println("Error:", err)
		return list, fmt.Errorf("can not open this file %s", err.Error())
	}
	defer file.Close()

	// 创建一个新的reader
	reader := csv.NewReader(file)

	data := make([]Record, 0, 50000)

	// 循环读取文件的每一行
	for {
		record, err := reader.Read()
		if err != nil {
			// 如果读取到文件的末尾，就跳出循环
			if err == io.EOF {
				break
			}
			// 如果遇到错误，就跳过这一行
			//fmt.Println("read csv file Error:", err)
			continue
		}

		// 检查字段数量
		if len(record) < 8 {
			//fmt.Println("Skipping line: not enough fields")
			continue
		}

		// 尝试解析这一行的数据
		alat, _ := strconv.ParseFloat(record[1], 64)
		alon, _ := strconv.ParseFloat(record[2], 64)
		lat, _ := strconv.ParseFloat(record[3], 64)
		lon, _ := strconv.ParseFloat(record[4], 64)
		cnt, _ := strconv.Atoi(record[5])
		rgn, _ := strconv.Atoi(record[6])
		city, _ := strconv.Atoi(record[7])

		data = append(data, Record{
			IP:   record[0],
			ALat: alat,
			ALon: alon,
			Lat:  lat,
			Lon:  lon,
			Cnt:  cnt,
			Rgn:  rgn,
			City: city,
		})
	}

	return data, nil
}
