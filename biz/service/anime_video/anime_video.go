package anime_video

import (
	"context"
	"eva/global"
	"eva/model/anime_video"
)

type AnimeService struct{}

func (a *AnimeService) CreateVideoSource(AnimeVideo anime_video.AnimeInfo) (uint, error) {
	ctx := context.Background()
	if err := global.EVA_DB.WithContext(ctx).Create(&AnimeVideo).Error; err != nil {
		return 0, err
	}
	return AnimeVideo.ID, nil
	//id, err := store.AnimeStore.Create(ctx, AnimeVideo)
	//if err != nil {
	//	return 0, err
	//}
	//return id, nil
}
