package utils

import "strconv"

func IntToBinaryArray(num int64, length int) []int {
	// 将十进制数转换为二进制字符串
	binaryStr := strconv.FormatInt(num, 2)

	// 填充0使其长度为指定长度
	for len(binaryStr) < length {
		binaryStr = "0" + binaryStr
	}

	// 将二进制字符串转换为二进制数组
	binaryArray := make([]int, length)
	for i, char := range binaryStr {
		if char == '1' {
			binaryArray[i] = 1
		} else {
			binaryArray[i] = 0
		}
	}

	return binaryArray
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

// toArrFromBitInt converts a binary integer to an array of integers.
func ToArrFromBitInt(num int64) []int {
	if num == 0 {
		return []int{0}
	}

	var arr []int
	for num > 0 {
		arr = append([]int{int(num % 2)}, arr...)
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
