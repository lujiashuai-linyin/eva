package store

import (
	"eva/biz/dal/store/anime_video"
	"eva/biz/dal/store/private"
)

var (
	CalendarStore private.CalendarStoreIf
	AnimeStore    anime_video.AnimeStoreIf
)
