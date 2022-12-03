package private

import (
	"eva/biz/service"
)

type ApiGroup struct {
	CalendarApi
}

var (
	CalendarService = service.ServiceGroupApp.PrivateServiceGroup.CalendarService
)
var (
	CalendarRpcService = service.ServiceGroupApp.PrivateServiceGroup.CalendarJobService
	jobService         = service.ServiceGroupApp.PrivateServiceGroup.JobService
)
