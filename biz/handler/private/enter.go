package private

import "eva/biz/service"

type RpcGroup struct {
	PrivateRpc
}

var (
	calendarService = service.ServiceGroupApp.CalendarServiceGroup.CalendarService
)
