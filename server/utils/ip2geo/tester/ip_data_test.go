package tester

import (
	"github.com/flipped-aurora/gin-vue-admin/server/utils/ip2geo/csv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdxIpData(t *testing.T) {
	list, err := csv.GetAdxIpList(csv.Ipv4, 1000)
	t.Logf("%#v", list)
	assert.Equal(t, err, nil)
}
