package private

import (
	"eva/biz/utils/timer"
	"eva/global"
	"eva/model/private"
	"fmt"
	"go.uber.org/zap"
)

type CalendarJobService struct{}

// 创建日历事件定时调度

func (j *CalendarJobService) Create(req *private.CalendarJob) (uint, error) {
	// 判断时间是否在本次轮询时间段内
	sendTimer := req.NoticeTime
	diff := sendTimer.Sub(timer.GetCurrTime())
	if diff < 0 {
		return 0, fmt.Errorf("您设置的提醒时间已经过了哦。")
	}
	//fmt.Println("hours:", diff.Hours())
	if diff.Hours() < 1 {
		//小于1个小时直接加入到定时器
		if err := global.EVA_DB.Create(&req).Error; err != nil {
			global.EVA_LOG.Error("内部错误！", zap.Error(err))
			return 0, err
		}
		go func() {
			jobService.HandleJob(*req)
		}()
		return req.ID, nil
	}
	req.Status = uint32(private.JobWait)
	if err := global.EVA_DB.Create(&req).Error; err != nil {
		global.EVA_LOG.Error("内部错误！", zap.Error(err))
		return 0, err
	}
	return req.ID, nil
}

func (j *CalendarJobService) GetJobsByTime(startTime string, endTime string) (jobs []private.CalendarJob, err error) {
	err = global.EVA_DB.Debug().Where("status=? and notice_time>=? and notice_time<=?", private.JobWait, startTime, endTime).
		Find(&jobs).Error
	return
}

func (j *CalendarJobService) UpdateStatusById(id, status int) error {
	return global.EVA_DB.Debug().Model(&private.CalendarJob{}).Where("id=?", id).Update("status", status).Error
}
