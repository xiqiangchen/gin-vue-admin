package utils

import "testing"

func TestToBitInt(t *testing.T) {
	st := []int{1, 3, 4, 9}
	t.Log(ToBitInt(st))
}
