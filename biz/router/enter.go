package router

import (
	"eva/biz/router/ai"
	"eva/biz/router/anime_video"
	"eva/biz/router/example"
	"eva/biz/router/private"
	"eva/biz/router/system"
)

type RouterGroup struct {
	System     system.RouterGroup
	Example    example.RouterGroup
	Private    private.RouterGroup
	ChatGpt    ai.RouterGroup
	AnimeVideo anime_video.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
