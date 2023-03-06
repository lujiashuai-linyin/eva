package ai

import "eva/biz/service"

type ApiGroup struct {
	ChatApi
}

var (
	chatService = service.ServiceGroupApp.AIServiceGroup.ChatService
)
