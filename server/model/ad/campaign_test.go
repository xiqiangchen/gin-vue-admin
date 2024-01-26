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
