package utils

import "math"

func Ceil(num float64, n int) float64 {
	pow := math.Pow(10, float64(n))
	return math.Ceil(num*pow) / pow
}
