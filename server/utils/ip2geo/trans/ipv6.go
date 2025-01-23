package trans

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"strconv"
	"strings"
	"sync"
)

type IpInfoV6 struct {
	prefStart map[uint32]uint32
	prefEnd   map[uint32]uint32
	data      []byte
	numbers   uint32
}

var objV6 *IpInfoV6
var lock sync.Mutex

func GetDistrictObjectV6(filePath string) (*IpInfoV6, error) {
	lock.Lock()
	defer lock.Unlock()
	var err error
	objV6, err = LoadFileV6(filePath)
	if err != nil {
		return nil, err
	}
	return objV6, nil
}

func LoadFileV6(file string) (*IpInfoV6, error) {
	p := IpInfoV6{}
	var err error
	p.data, err = ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	p.numbers = unpackInt4byteV6(p.data[4], p.data[5], p.data[6], p.data[7])
	p.prefStart = make(map[uint32]uint32)
	p.prefEnd = make(map[uint32]uint32)
	for k := uint32(0); k < p.numbers; k++ {
		i := k*12 + 4 + 4
		p.prefStart[unpackInt4byteV6(p.data[i+8], p.data[i+9], p.data[i+10], p.data[i+11])] = unpackInt4byteV6(p.data[i], p.data[i+1], p.data[i+2], p.data[i+3])
		p.prefEnd[unpackInt4byteV6(p.data[i+8], p.data[i+9], p.data[i+10], p.data[i+11])] = unpackInt4byteV6(p.data[i+4], p.data[i+5], p.data[i+6], p.data[i+7])
	}

	return &p, err

}

func (p *IpInfoV6) Get(ip string) (string, error) {
	ips := strings.Split(ip, ":")

	if len(ips) == 0 {
		return "", errors.New("IP address is empty")
	}

	parseUint, err := strconv.ParseUint(ips[0], 16, 32)
	if err != nil {
		return "", fmt.Errorf("failed to parse IP address: %v", err)
	}

	prefix := uint32(parseUint)

	intIP, err := ipToIntV6(ip)
	if err != nil {
		return "", err
	}

	low := p.prefStart[prefix]
	high := p.prefEnd[prefix]

	var cur uint32
	if low == high {
		cur = low
	} else {
		cur = p.SearchV6(low, high, intIP)
	}
	return p.getAddr(cur)
}

func (p *IpInfoV6) getAddr(cur uint32) (string, error) {
	j := p.numbers*12 + 4 + 4 + (cur * 55)
	if int(j)+54 >= len(p.data) {
		return "", errors.New("index out of range")
	}

	offset := unpackInt4byteV6(p.data[j+50], p.data[1+j+50], p.data[2+j+50], p.data[3+j+50])
	length := uint32(p.data[50+j+4])

	if int(offset+length) > len(p.data) {
		return "", errors.New("index out of range")
	}

	return string(p.data[offset:int(offset+length)]), nil
}

func (p *IpInfoV6) SearchV6(low uint32, high uint32, k *big.Int) uint32 {
	var M uint32 = 0
	for low <= high {
		mid := (low + high) / 2
		j := p.numbers*12 + 4 + 4 + (mid * 55)
		endipnum := new(big.Int)
		endipnumInt, _ := endipnum.SetString(strings.ReplaceAll(string(p.data[j:j+50]), "*", ""), 10)

		if endipnumInt.Cmp(k) == 0 || endipnumInt.Cmp(k) == 1 {
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

func ipToIntV6(ipstr string) (*big.Int, error) {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return nil, fmt.Errorf("invalid IP address: %s", ipstr)
	}
	ip = ip.To16()
	IPv6Int := big.NewInt(0)
	IPv6Int.SetBytes(ip)
	return IPv6Int, nil
}

func unpackInt4byteV6(a, b, c, d byte) uint32 {
	return (uint32(a) & 0xFF) | ((uint32(b) << 8) & 0xFF00) | ((uint32(c) << 16) & 0xFF0000) | ((uint32(d) << 24) & 0xFF000000)
}
