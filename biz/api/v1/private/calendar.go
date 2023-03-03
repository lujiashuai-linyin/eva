package private

import (
	"eva/biz/utils/timer"
	"eva/global"
	"eva/model/common/response"
	"eva/model/private"
	"eva/model/system/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
	"time"
)

type CalendarApi struct{}

// 创建事件，添加定时任务

func (b *CalendarApi) AddCalendarRecord(c *gin.Context) {
	req := private.Calendar{}
	err := c.ShouldBindJSON(&req)
	jobIDLIst := []uint{}
	claims, _ := c.Get("claims")
	userID := claims.(*request.CustomClaims).ID
	req.UserID = int(userID)
	if err != nil {
		global.EVA_LOG.Error("请求参数错误!", zap.Error(err))
		response.FailWithMessage("请求参数错误！", c)
		return
	}
	// 创建记录

	recordID, err := CalendarService.CreateCalendarRecord(req)
	if err != nil {
		global.EVA_LOG.Error("内部错误!", zap.Error(err))
		response.SFailWithMessage("内部错误！", c)
		return
	}
	// 创建添加定时任务
	if req.Alert {
		alertList := strings.Split(req.PreAlertTime, ",")
		for _, alertTime := range alertList {
			pre, _ := time.ParseDuration("-" + alertTime + "m")
			job := private.CalendarJob{
				UserID:     uint64(req.UserID),
				RecordID:   uint64(recordID),
				Title:      req.Title,
				Content:    timer.TimeString(req.StartTime) + "——" + timer.TimeString(req.EndTime) + "\n" + req.Content,
				NoticeTime: req.StartTime.Add(pre),
				Email:      "690067698@qq.com",
			}
			jobID, err := CalendarRpcService.Create(&job)
			if err != nil {
				global.EVA_LOG.Error("内部错误!", zap.Error(err))
				response.SFailWithMessage("内部错误！", c)
				return
			}
			jobIDLIst = append(jobIDLIst, jobID)
		}
	}
	// TODO
	// 更新record表，把jobID加进去
	fmt.Println("jobIDLIst", jobIDLIst)
	if err != nil {
		global.EVA_LOG.Error("内部错误!", zap.Error(err))
		response.SFailWithMessage("内部错误！", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// 获取事件列表

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

//删除事件与定时任务，指事件也删除

func (b *CalendarApi) DeleteCalendarRecord(c *gin.Context) {
	res := make(map[string]interface{})
	err := c.ShouldBindJSON(&res)
	if err != nil {
		global.EVA_LOG.Error("请求参数错误!", zap.Error(err))
		response.FailWithMessage("请求参数错误！", c)
		return
	}
	claims, _ := c.Get("claims")
	userID := claims.(*request.CustomClaims).ID
	err = CalendarService.DeleteCalendarRecord(userID, int(res["id"].(float64)))
	if err != nil {
		global.EVA_LOG.Error("内部错误!", zap.Error(err))
		response.SFailWithMessage("内部错误！", c)
		return
	}
	response.OkWithMessage("删除成功", c)
	_ = CalendarRpcService.DeleteAndRemoveJob(int(res["id"].(float64)))
}

// 启动事件提醒服务，仅管理员有权限

func (b *CalendarApi) JobServerStart(c *gin.Context) {
	claims, _ := c.Get("claims")
	authorityID := claims.(*request.CustomClaims).AuthorityId
	if authorityID != 1 {
		global.EVA_LOG.Error("权限不足!")
		response.FailWithMessage("权限不足！", c)
		return
	}
	jobService.SchedulerCalendarInit()
}

// 更新事件,与定时任务，包括: 将不提醒设置成提醒，删除或新增部分提醒时间, 改动开始时间等

func (b *CalendarApi) UpdateCalendarRecord(c *gin.Context) {
	req := private.Calendar{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.EVA_LOG.Error("请求参数错误!", zap.Error(err))
		response.FailWithMessage("请求参数错误！", c)
		return
	}
	claims, _ := c.Get("claims")
	userID := claims.(*request.CustomClaims).ID
	recordOld, err := CalendarService.GetCalendarRecordDetail(userID, int(req.ID))
	err = CalendarService.UpdateCalendarRecord(userID, req)
	if err != nil {
		global.EVA_LOG.Error("内部错误!", zap.Error(err))
		response.FailWithMessage("内部错误！", c)
		return
	}
	req.UserID = int(userID)
	err = CalendarRpcService.UpdateCronAndJob(recordOld, req)
	if err != nil {
		global.EVA_LOG.Error("内部错误!", zap.Error(err))
		response.FailWithMessage("内部错误！", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// 更新结束时间

func (b *CalendarApi) UpdateEndTime(c *gin.Context) {
	req := private.Calendar{}
	err := c.ShouldBindJSON(&req)
	fmt.Println(req)
	if err != nil {
		global.EVA_LOG.Error("请求参数错误!", zap.Error(err))
		response.FailWithMessage("请求参数错误！", c)
		return
	}
	claims, _ := c.Get("claims")
	userID := claims.(*request.CustomClaims).ID
	recordOld, err := CalendarService.GetCalendarRecordDetail(userID, int(req.ID))
	err = CalendarService.UpdateEndTime(userID, req)
	if err != nil {
		global.EVA_LOG.Error("内部错误!", zap.Error(err))
		response.FailWithMessage("内部错误！", c)
		return
	}
	recordNew := recordOld
	recordNew.StartTime = req.StartTime
	recordNew.EndTime = req.EndTime
	err = CalendarRpcService.UpdateCronAndJob(recordOld, recordNew)
	if err != nil {
		global.EVA_LOG.Error("内部错误!", zap.Error(err))
		response.FailWithMessage("内部错误！", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}
