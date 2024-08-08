package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/ad"
	"github.com/flipped-aurora/gin-vue-admin/server/service/assert"
	"github.com/flipped-aurora/gin-vue-admin/server/service/dsp"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/resource"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	ResourceServiceGroup resource.ServiceGroup
	AssertServiceGroup   assert.ServiceGroup
	AdServiceGroup       ad.ServiceGroup
	DspGroup             dsp.ServiceGroup
}
