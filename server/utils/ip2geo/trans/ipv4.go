package trans

import (
	"encoding/binary"
	"errors"
	"io/ioutil"
	"net"
	"strconv"
	"strings"
)

type IpInfo struct {
	prefStart [256]uint32
	prefEnd   [256]uint32
	data      []byte
}

var obj *IpInfo

func GetObject(fileName string) (*IpInfo, error) {
	obj = &IpInfo{}
	var err error
	obj, err = LoadFile(fileName)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func LoadFile(file string) (*IpInfo, error) {
	p := IpInfo{}
	data, err := ioutil.ReadFile(file)
	p.data = data
	if err != nil {
		return nil, err
	}

	for k := 0; k < 256; k++ {
		i := k*8 + 4
		p.prefStart[k] = unpackInt4byte(data[i], data[i+1], data[i+2], data[i+3])
		p.prefEnd[k] = unpackInt4byte(data[i+4], data[i+5], data[i+6], data[i+7])
	}

	return &p, err

}

func (p *IpInfo) Get(ip string) (string, error) {
	ips := strings.Split(ip, ".")
	x, _ := strconv.Atoi(ips[0])
	prefix := uint32(x)
	intIP, err := ipToInt(ip)
	if err != nil {
		return "", err
	}

	low := p.prefStart[prefix]
	high := p.prefEnd[prefix]

	var cur uint32
	if low == high {
		cur = low
	} else {
		cur = p.Search(low, high, intIP)
	}
	if cur == 100000000 {
		return "", errors.New("无信息")
	} else {
		return p.getAddr(cur), nil
	}

}

func (p *IpInfo) Search(low uint32, high uint32, k uint32) uint32 {
	var M uint32 = 0
	for low <= high {
		mid := (low + high) / 2

		j := 2052 + (mid * 9)
		endipNum := unpackInt4byte(p.data[j], p.data[1+j], p.data[2+j], p.data[3+j])

		if endipNum >= k {
			M = mid
			if mid == 0 {
				break
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return M
}

func (p *IpInfo) getAddr(cur uint32) string {
	j := 2052 + (cur * 9)
	offset := unpackInt4byte(p.data[4+j], p.data[5+j], p.data[6+j], p.data[7+j])
	length := uint32(p.data[8+j])
	return string(p.data[offset:int(offset+length)])
}

func ipToInt(ipstr string) (uint32, error) {
	ip := net.ParseIP(ipstr)
	ip = ip.To4()
	if ip == nil {
		return 0, errors.New("ip 不合法")
	}
	return binary.BigEndian.Uint32(ip), nil
}

func unpackInt4byte(a, b, c, d byte) uint32 {
	return (uint32(a) & 0xFF) | ((uint32(b) << 8) & 0xFF00) | ((uint32(c) << 16) & 0xFF0000) | ((uint32(d) << 24) & 0xFF000000)
}
