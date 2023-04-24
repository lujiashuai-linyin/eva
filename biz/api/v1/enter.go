package v1

import (
	"eva/biz/api/v1/ai"
	"eva/biz/api/v1/anime_video"
	"eva/biz/api/v1/example"
	"eva/biz/api/v1/private"
	"eva/biz/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	FileApiGroup    example.ApiGroup
	PrivateApiGroup private.ApiGroup
	ChatApiGroup    ai.ApiGroup
	AnimeApiGroup   anime_video.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
