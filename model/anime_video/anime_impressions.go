package anime_video

import "eva/global"

// redis

type AnimeImpressions struct {
	global.EVA_MODEL
	AnimeId uint `json:"anime_id" gorm:"column: anime_id;comment:动漫id"`
	// 订阅数量
	// 收藏数量
	// 点赞数量
	Extra string `json:"extra" form:"extra" gorm:"column:extra;comment:extra"`
}

func (AnimeImpressions) TableName() string {
	return "anime_impressions"
}
