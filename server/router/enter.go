package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/ad"
	"github.com/flipped-aurora/gin-vue-admin/server/router/assert"
	"github.com/flipped-aurora/gin-vue-admin/server/router/dsp"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/resource"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Resource resource.RouterGroup
	Assert   assert.RouterGroup
	Ad       ad.RouterGroup
	Dsp      dsp.RouterGroup
}
