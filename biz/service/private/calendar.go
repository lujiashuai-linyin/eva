package private

import (
	"context"
	"eva/biz/dal/store"
	"eva/model/private"
	"fmt"
	"time"
)

type CalendarService struct{}

var calendarJobService = CalendarJobService{}
var jobService = JobService{}

// 创建日历事件

func (c *CalendarService) CreateCalendarRecord(record private.Calendar) (uint, error) {
	ctx := context.Background()
	var realID, id uint
	oneDay, err := time.ParseDuration("24h")
	switch record.Repeat {
	case "每天":
		i := 0
		for ; i < 7; i++ {
			id, err = store.CalendarStore.Create(ctx, record)
			if i == 0 {
				fmt.Println("realID:", id)
				realID = id
			}
			if err != nil {
				break
			}
			record.StartTime = record.StartTime.Add(oneDay)
			record.EndTime = record.EndTime.Add(oneDay)
		}

		if err != nil {
			return 0, err
		}
	case "每周":
		i := 0
		for ; i < 7; i++ {
			id, err = store.CalendarStore.Create(ctx, record)
			if i == 0 {
				fmt.Println("realID:", id)
				realID = id
			}
			if err != nil {
				break
			}
			record.StartTime = record.StartTime.AddDate(0, 0, 7)
			record.EndTime = record.EndTime.AddDate(0, 0, 7)
		}
		if err != nil {
			return 0, err
		}
	case "每月":
		i := 0
		year, month, day := record.StartTime.Date()
		id, err = store.CalendarStore.Create(ctx, record)
		realID = id
		if err != nil {
			break
		}
		for ; i < 6; i++ {
			yearGo, monthGo, dayGo := record.StartTime.AddDate(0, 1, 0).Date()
			if yearGo == year && monthGo == month+1 && dayGo == day {
				record.StartTime = record.StartTime.AddDate(0, 1, 0)
				record.EndTime = record.EndTime.AddDate(0, 1, 0)
				id, err = store.CalendarStore.Create(ctx, record)
			} else if yearGo == year+1 && month == 12 && dayGo == day {
				record.StartTime = record.StartTime.AddDate(0, 1, 0)
				record.EndTime = record.EndTime.AddDate(0, 1, 0)
				id, err = store.CalendarStore.Create(ctx, record)
			}
		}
		if err != nil {
			return 0, err
		}
	case "每年":
		i := 0
		for ; i < 7; i++ {
			id, err = store.CalendarStore.Create(ctx, record)
			if i == 0 {
				fmt.Println("realID:", id)
				realID = id
			}
			record.StartTime = record.StartTime.AddDate(0, 0, 1)
			record.EndTime = record.EndTime.AddDate(0, 0, 1)
		}
		if err != nil {
			return 0, err
		}
	default:
		id, err = store.CalendarStore.Create(ctx, record)
		realID = id
		if err != nil {
			return 0, err
		}
	}
	return realID, nil
}

// 获取日历事件

func (c *CalendarService) GetCalendarRecord(id uint) ([]*private.Calendar, error) {
	ctx := context.Background()
	result, err := store.CalendarStore.FindRecordList(ctx, int(id))
	if err != nil {
		return result, err
	}
	return result, nil
}

func (c *CalendarService) DeleteCalendarRecord(id int) error {
	ctx := context.Background()
	err := store.CalendarStore.DeleteRecord(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
