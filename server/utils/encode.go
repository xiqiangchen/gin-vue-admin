package utils

import (
	"hash/fnv"
	"strconv"
)

func HashCode(s string) uint64 {
	h := fnv.New64()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}

// findMax finds the maximum integer in an array.
func findMax(arr []int) int {
	max := arr[0]
	for _, value := range arr[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

func GetId(val string) uint64 {
	var ret uint64
	if len(val) > 0 {
		if puer, err := strconv.Atoi(val); err != nil {
			ret = HashCode(val)
		} else {
			ret = uint64(puer)
		}
	}
	return ret
}
