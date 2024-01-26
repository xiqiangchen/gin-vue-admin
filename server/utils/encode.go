package utils

import "hash/fnv"

func HashCode(s string) uint64 {
	h := fnv.New64()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}

// toBitInt converts an array of integers to a binary integer.
func ToBitInt(arr []int) int {
	max := findMax(arr)
	newArr := make([]int, max+1)

	for _, value := range arr {
		newArr[value] = 1
	}

	num := 0
	for _, value := range newArr {
		num = num << 1
		num = num + value
	}

	return num
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

// toArrFromBitInt converts a binary integer to an array of integers.
func ToArrFromBitInt(num int) []int {
	if num == 0 {
		return []int{0}
	}

	var arr []int
	for num > 0 {
		arr = append([]int{num % 2}, arr...)
		num = num / 2
	}

	var newArr []int
	for i, bit := range arr {
		if bit != 0 {
			newArr = append(newArr, i+1)
		}
	}

	return newArr
}
