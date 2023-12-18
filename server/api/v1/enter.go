package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/assert"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/resource"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	ResourceApiGroup resource.ApiGroup
	AssertApiGroup   assert.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
