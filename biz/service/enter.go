package service

import (
	"eva/biz/service/ai"
	"eva/biz/service/example"
	"eva/biz/service/private"
	"eva/biz/service/send_message"
	"eva/biz/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	PrivateServiceGroup private.ServiceGroup
	SendMessageService  send_message.ServiceGroup
	AIServiceGroup      ai.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
