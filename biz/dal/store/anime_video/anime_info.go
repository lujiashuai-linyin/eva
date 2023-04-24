package anime_video

import (
	"context"
	"eva/global"
	"eva/model/anime_video"
)

type AnimeStoreIf interface {
	//FindRecordList(ctx context.Context, UserID int) ([]*private.Calendar, error)
	//DeleteRecord(ctx context.Context, UserID uint, RecordID int) error
	Create(ctx context.Context, AnimeInfo anime_video.AnimeInfo) (uint, error)
	//Update(ctx context.Context, UserID uint, Record private.Calendar) error
	//Updata(ctx context.Context, UserID uint, Record private.Calendar) error
	//Detail(ctx context.Context, AnimeID uint) (anime_video.AnimeInfo, error)
}

type AnimeStore struct{}

// 创建一条事件

func (store *AnimeStore) Create(ctx context.Context, AnimeInfo anime_video.AnimeInfo) (uint, error) {
	if err := global.EVA_DB.WithContext(ctx).Debug().Create(&AnimeInfo).Error; err != nil {
		return 0, err
	}
	return AnimeInfo.ID, nil
}
