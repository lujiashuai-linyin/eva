package private

import (
	"eva/biz/handler"
	"eva/biz/service"
)

type ApiGroup struct {
	CalendarApi
}

var (
	CalendarService = service.ServiceGroupApp.CalendarServiceGroup.CalendarService
)
var (
	CalendarRpcService = handler.RpcGroupApp.BytestRpcGroup.PrivateRpc
)
