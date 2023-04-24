package anime_video

import (
	v1 "eva/biz/api/v1"
	"github.com/gin-gonic/gin"
)

type AnimeRouter struct{}

func (s *AnimeRouter) InitAnimeRouter(Router *gin.RouterGroup) {
	gptRouterWithJWT := Router.Group("anime").Use()
	animeApi := v1.ApiGroupApp.AnimeApiGroup.AnimeApi
	{
		gptRouterWithJWT.POST("create", animeApi.CreateAnimeVideo)
		//gptRouterWithJWT.GET("topic_chat", chatApi.GetTopicMessage)
		//gptRouterWithJWT.GET("topic_list", chatApi.GetTopicList)
	}
}
