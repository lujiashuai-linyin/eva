package anime_video

import (
	"eva/global"
	"eva/model/anime_video"
	"eva/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AnimeApi struct{}

func (a *AnimeApi) CreateAnimeVideo(c *gin.Context) {
	req := anime_video.AnimeInfo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.EVA_LOG.Error("请求参数错误!", zap.Error(err))
		response.FailWithMessage("请求参数错误！", c)
		return
	}

	id, err := animeService.CreateVideoSource(req)
	if err != nil {
		global.EVA_LOG.Error("内部错误!", zap.Error(err))
		response.SFailWithMessage("内部错误！", c)
		return
	}
	response.OkWithDetailed(map[string]interface{}{"id": id}, "创建成功", c)
}
