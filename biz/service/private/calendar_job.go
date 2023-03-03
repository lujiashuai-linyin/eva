package private

import (
	"eva/biz/utils/timer"
	"eva/global"
	"eva/model/private"
	"fmt"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"strings"
	"time"
)

type CalendarJobService struct{}

// 业务
// 删除Job和Cron
func (j *CalendarJobService) DeleteAndRemoveJob(record_id int) error {
	jobList, err := j.GetCronIDByRecordID(record_id)
	if err != nil {
		return err
	}
	err = j.DeleteJobByRecordID(record_id)
	if err != nil {
		return err
	}
	for _, item := range jobList {
		global.EVA_Timer.Remove("calendarSendAlert", int(item.CronID))
	}
	return nil
}

// 更新Job和Cron
func (j *CalendarJobService) UpdateCronAndJob(oldRecord, newRecord private.Calendar) error {
	// 两次都是不提醒
	if !newRecord.Alert && !oldRecord.Alert {
		return nil
	}
	// 判断是否更改提醒
	if newRecord.Alert && !oldRecord.Alert {
		alertList := strings.Split(newRecord.PreAlertTime, ",")
		jobIDLIst := []uint{}
		for _, alertTime := range alertList {
			pre, _ := time.ParseDuration("-" + alertTime + "m")
			job := private.CalendarJob{
				UserID:     uint64(newRecord.UserID),
				RecordID:   uint64(newRecord.ID),
				Title:      newRecord.Title,
				Content:    timer.TimeString(newRecord.StartTime) + "——" + timer.TimeString(newRecord.EndTime) + "\n" + newRecord.Content,
				NoticeTime: newRecord.StartTime.Add(pre),
				Email:      "690067698@qq.com",
			}
			jobID, err := j.Create(&job)
			if err != nil {
				return err
			}
			jobIDLIst = append(jobIDLIst, jobID)
		}
		fmt.Println(jobIDLIst)
		return nil
	}
	if oldRecord.Alert && !newRecord.Alert {
		err := j.DeleteAndRemoveJob(int(newRecord.ID))
		fmt.Println(err)
		return err
	}
	// 如果两次都是提醒
	if oldRecord.Alert && newRecord.Alert {
		// 更改了提醒列表, 采用先删后加的逻辑，这样实现起来最容易
		err := j.DeleteAndRemoveJob(int(newRecord.ID))
		if err != nil {
			fmt.Println(err)
			return err
		}
		alertList := strings.Split(newRecord.PreAlertTime, ",")
		jobIDLIst := []uint{}
		for _, alertTime := range alertList {
			pre, _ := time.ParseDuration("-" + alertTime + "m")
			job := private.CalendarJob{
				UserID:     uint64(newRecord.UserID),
				RecordID:   uint64(newRecord.ID),
				Title:      newRecord.Title,
				Content:    timer.TimeString(newRecord.StartTime) + "——" + timer.TimeString(newRecord.EndTime) + "\n" + newRecord.Content,
				NoticeTime: newRecord.StartTime.Add(pre),
				Email:      "690067698@qq.com",
			}
			jobID, err := j.Create(&job)
			if err != nil {
				return err
			}
			jobIDLIst = append(jobIDLIst, jobID)
		}
		fmt.Println(jobIDLIst)
		return nil
	}
	return nil
}

// 数据库操作

func (j *CalendarJobService) GetJobsByTime(startTime string, endTime string) (jobs []private.CalendarJob, err error) {
	err = global.EVA_DB.Debug().Where("status=? and notice_time>=? and notice_time<=?", private.JobWait, startTime, endTime).
		Find(&jobs).Error
	return
}

func (j *CalendarJobService) Create(req *private.CalendarJob) (uint, error) {
	// 判断时间是否在本次轮询时间段内
	sendTimer := req.NoticeTime
	diff := sendTimer.Sub(timer.GetCurrTime())
	if diff < 0 {
		return 0, fmt.Errorf("您设置的提醒时间已经过了哦。")
	}
	fmt.Println("hours:", diff.Hours())
	if diff.Hours() < 1 {
		//小于1个小时直接加入到定时器
		if err := global.EVA_DB.Create(&req).Error; err != nil {
			global.EVA_LOG.Error("内部错误！", zap.Error(err))
			return 0, err
		}
		jobService.HandleJob(*req)
		return req.ID, nil
	}
	req.Status = uint32(private.JobWait)
	if err := global.EVA_DB.Create(&req).Error; err != nil {
		global.EVA_LOG.Error("内部错误！", zap.Error(err))
		return 0, err
	}
	return req.ID, nil
}

func (j *CalendarJobService) GetCronID(id int) (uint, error) {
	var job private.CalendarJob
	err := global.EVA_DB.Debug().Where("id=?", id).Select("cron_id").First(&job).Error
	if err != nil {
		return 0, err
	}
	return job.ID, nil
}

func (j *CalendarJobService) DeleteJobByRecordID(id int) error {
	if err := global.EVA_DB.Debug().Where("record_id = ?", id).Delete(&private.CalendarJob{}).Error; err != nil {
		return err
	}
	return nil
}

func (j *CalendarJobService) GetCronIDByRecordID(id int) ([]private.CalendarJob, error) {
	var job []private.CalendarJob
	err := global.EVA_DB.Debug().Where("record_id=?", id).Select("cron_id").First(&job).Error
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (j *CalendarJobService) UpdateStatusById(id, status int) error {
	return global.EVA_DB.Debug().Model(&private.CalendarJob{}).Where("id=?", id).Update("status", status).Error
}

func (j *CalendarJobService) UpdateCronIDAndStatusById(id int, cronID cron.EntryID, status int) error {
	return global.EVA_DB.Debug().Model(&private.CalendarJob{}).Where("id=?", id).Updates(map[string]interface{}{"cron_id": uint64(cronID), "status": status}).Error
}
