package utils

import "hash/fnv"

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
