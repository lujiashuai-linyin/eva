package private

import (
	"eva/global"
	"time"
)

type Calendar struct {
	global.EVA_MODEL
	IsAllDay     bool      `json:"is_all_day" form:"is_all_day" gorm:"column:is_all_day;comment:是否全天"`
	RecordType   string    `json:"record_type" form:"record_type" gorm:"column:record_type;type:varchar(4);comment:日程类型"`
	StartTime    time.Time `json:"start_time" form:"start_time" gorm:"column:start_time;comment:开始时间"`
	EndTime      time.Time `json:"end_time" form:"end_time" gorm:"column:end_time;comment:结束时间"`
	Title        string    `json:"title"  form:"title" gorm:"column:title;type:varchar(200);comment:title"`
	Content      string    `json:"content" form:"content" gorm:"column:content;type:text;comment:content"`
	Alert        bool      `json:"alert" form:"alert" gorm:"column:alert;comment:是否提醒"`
	PreAlertTime string    `json:"pre_alert_time" form:"pre_alert_time" gorm:"column:pre_alert_time;type:varchar(18);comment:提前多久提醒,为列表"`
	Address      string    `json:"address" form:"address" gorm:"column:address;type:varchar(500);comment:文件地址"`
	UserStatus   string    `json:"user_status" form:"user_status" gorm:"column:user_status;type:varchar(5);comment:设置用户状态"`
	Repeat       string    `json:"repeat" form:"repeat" gorm:"column:repeat;type:varchar(10);comment:重复策略/每天/每周/每月/每年"`
	UserID       int       `json:"user_id" form:"user_id" gorm:"column:user_id;"`
	Place        string    `json:"place" form:"place" gorm:"column:place;type:varchar(100);comment:地点"`
	Img          string    `json:"img" form:"img" gorm:"column:img;type:varchar(500);comment:图片地址"`
	Extra        string    `json:"extra" form:"extra" gorm:"column:extra;comment:extra"`
}

func (Calendar) TableName() string {
	return "calendars"
}
