package anime_video

import "eva/biz/service"

type ApiGroup struct {
	AnimeApi
}

var (
	animeService = service.ServiceGroupApp.AnimeServiceGroup.AnimeService
)
