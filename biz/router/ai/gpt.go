package ai

import (
	v1 "eva/biz/api/v1"
	"eva/middleware"
	"github.com/gin-gonic/gin"
)

type ChatRouter struct{}

func (s *ChatRouter) InitGptRouter(Router *gin.RouterGroup) {
	gptRouterWithJWT := Router.Group("gpt").Use(middleware.JWTAuth())
	chatApi := v1.ApiGroupApp.ChatApiGroup.ChatApi
	{
		gptRouterWithJWT.POST("question", chatApi.Question)
		gptRouterWithJWT.POST("topic_delete", chatApi.DeleteTopic)
		gptRouterWithJWT.GET("topic_chat", chatApi.GetTopicMessage)
		gptRouterWithJWT.GET("topic_list", chatApi.GetTopicList)
	}
}
