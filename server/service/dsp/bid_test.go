package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"math/rand"
	"testing"
)

func TestBid(t *testing.T) {
	bidPrice := 0.82
	bidFloor := 0.67
	for i := 0; i < 100; i++ {
		t.Log(utils.Ceil(((bidPrice-bidFloor)*0.3+bidFloor)+0.35*rand.Float64()*(bidPrice-bidFloor), 2))
	}
}
