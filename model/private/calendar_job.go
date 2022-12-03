package private

import (
	"eva/global"
	"time"
)

var (
	// 通知成功
	JobSuccess = 1
	// 待通知
	JobWait = 2
	// 通知失败
	JobFail = 3
	// 等候中
	JobDoing = 4
)

type CalendarJob struct {
	global.EVA_MODEL
	UserID     uint64    `json:"user_id" form:"user_id" gorm:"column:user_id;"`
	RecordID   uint64    `json:"record_id" form:"record_id" gorm:"column:record_id;"`
	Status     uint32    `json:"status" form:"status" gorm:"column:status;type:tinyint;default:2"`
	Title      string    `json:"title"  form:"title" gorm:"column:title;type:varchar(200);comment:title"`
	Content    string    `json:"content" form:"content" gorm:"column:content;type:text;comment:content"`
	NoticeTime time.Time `json:"notice_time" form:"notice_time" gorm:"column:notice_time;comment:提醒时间"`
	Phone      string    `json:"phone" form:"phone" gorm:"column:phone;type:varchar(13);comment:手机号"`
	Email      string    `json:"email" form:"email" gorm:"column:email;type:varchar(50);comment:邮箱"`
	WeChat     string    `json:"wechat" form:"wechat" gorm:"column:wechat;type:varchar(50);comment:微信号"`
}

func (CalendarJob) TableName() string {
	return "calendar_jobs"
}
