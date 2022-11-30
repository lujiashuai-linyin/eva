package router

import (
	"eva/biz/router/example"
	"eva/biz/router/private"
	"eva/biz/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Private private.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
