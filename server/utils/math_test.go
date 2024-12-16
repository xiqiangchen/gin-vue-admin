package utils

import (
	"testing"
)

func TestCeil2(t *testing.T) {
	t.Log(Ceil(123.4111567, 2))
	t.Log(Ceil(123.4567, 2))
	t.Log(Ceil(0.1234567, 4))
	t.Log(Ceil(123.4567, 0))
	t.Log(Ceil(123.4567, -1))
}

func TestWeightedRandomIndex(t *testing.T) {
	for i := 0; i < 100; i++ {
		weights := []float64{0.1, 0.2, 0.3, 0.4}
		index, _ := WeightedRandomIndex(weights)
		t.Logf("Selected index: %d\n", index)
	}
}
