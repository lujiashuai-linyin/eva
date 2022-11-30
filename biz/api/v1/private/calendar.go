package private

import (
	"eva/global"
	"eva/model/common/response"
	"eva/model/private"
	"eva/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CalendarApi struct{}

func (b *CalendarApi) AddCalendarRecord(c *gin.Context) {
	req := private.Calendar{}
	err := c.ShouldBindJSON(&req)
	claims, _ := c.Get("claims")
	userID := claims.(*request.CustomClaims).ID
	req.UserID = int(userID)
	if err != nil {
		global.EVA_LOG.Error("请求参数错误!", zap.Error(err))
		response.FailWithMessage("请求参数错误！", c)
		return
	}
	// 创建记录
	err = CalendarService.CreateCalendarRecord(req)
	if err != nil {
		global.EVA_LOG.Error("内部错误!", zap.Error(err))
		response.SFailWithMessage("内部错误！", c)
		return
	}
	// 创建添加定时任务
	err = CalendarService.CreateCalendarRecordScheduler(req)
	if err != nil {
		global.EVA_LOG.Error("内部错误!", zap.Error(err))
		response.SFailWithMessage("内部错误！", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (b *CalendarApi) GetCalendarRecord(c *gin.Context) {
	claims, _ := c.Get("claims")
	result, err := CalendarService.GetCalendarRecord(claims.(*request.CustomClaims).ID)
	if err != nil {
		global.EVA_LOG.Error("内部错误!", zap.Error(err))
		response.SFailWithMessage("内部错误！", c)
		return
	}
	response.OkWithDetailed(result, "获取成功", c)
}

func (b *CalendarApi) DeleteCalendarRecord(c *gin.Context) {
	res := make(map[string]interface{})
	err := c.ShouldBindJSON(&res)
	if err != nil {
		global.EVA_LOG.Error("请求参数错误!", zap.Error(err))
		response.FailWithMessage("请求参数错误！", c)
		return
	}
	err = CalendarService.DeleteCalendarRecord(int(res["id"].(float64)))
	if err != nil {
		global.EVA_LOG.Error("内部错误!", zap.Error(err))
		response.SFailWithMessage("内部错误！", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
