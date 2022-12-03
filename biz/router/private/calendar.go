package private

import (
	v1 "eva/biz/api/v1"
	"eva/middleware"
	"github.com/gin-gonic/gin"
)

type CalendarRouter struct{}

func (s *CalendarRouter) InitCalendarRouter(Router *gin.RouterGroup) {
	calendarRouterWithJWT := Router.Group("calendar").Use(middleware.JWTAuth())
	calendarApi := v1.ApiGroupApp.PrivateApiGroup.CalendarApi
	{
		calendarRouterWithJWT.POST("/add_record", calendarApi.AddCalendarRecord)
		calendarRouterWithJWT.GET("/get_record_list", calendarApi.GetCalendarRecord)
		calendarRouterWithJWT.POST("/delete_record", calendarApi.DeleteCalendarRecord)
		calendarRouterWithJWT.GET("/start_job_server", calendarApi.JobServerStart)
	}
}
