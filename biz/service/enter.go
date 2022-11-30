package service

import (
	"eva/biz/service/example"
	"eva/biz/service/private"
	"eva/biz/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	CalendarServiceGroup private.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
