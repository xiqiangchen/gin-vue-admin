package ad

import (
	"math/rand"
	"testing"
)

func TestGetBidPrice(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(((rand.Float64()-0.5)*0.2 + 1) * 10)
	}
}

func TestGetBidRate(t *testing.T) {
	var t1, t2 int
	for i := 0; i < 100; i++ {
		b := rand.Float64() > float64(10)/100
		if b {
			t1++
		} else {
			t2++
		}
	}
	t.Log(t1, t2)
}
