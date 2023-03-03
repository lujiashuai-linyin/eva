package private

import (
	"context"
	"eva/global"
	"eva/model/private"
)

type CalendarStoreIf interface {
	FindRecordList(ctx context.Context, UserID int) ([]*private.Calendar, error)
	DeleteRecord(ctx context.Context, UserID uint, RecordID int) error
	Create(ctx context.Context, Record private.Calendar) (uint, error)
	Update(ctx context.Context, UserID uint, Record private.Calendar) error
	UpdateEndTime(ctx context.Context, UserID uint, Record private.Calendar) error
	Detail(ctx context.Context, UserID uint, RecordID int) (private.Calendar, error)
}

type CalendarStore struct{}

// 创建一条事件

func (store *CalendarStore) Create(ctx context.Context, Record private.Calendar) (uint, error) {
	if err := global.EVA_DB.Create(&Record).Error; err != nil {
		return 0, err
	}
	return Record.ID, nil
}

// 获取事件列表

func (store *CalendarStore) FindRecordList(ctx context.Context, UserID int) ([]*private.Calendar, error) {
	result := []*private.Calendar{}
	if err := global.EVA_DB.WithContext(ctx).Where("user_id = ?", UserID).Find(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
func (store *CalendarStore) DeleteRecord(ctx context.Context, UserID uint, RecordID int) error {
	if err := global.EVA_DB.WithContext(ctx).Where("id = ? and user_id = ?", RecordID, UserID).Delete(&private.Calendar{}).Error; err != nil {
		return err
	}
	return nil
}
func (store *CalendarStore) Update(ctx context.Context, UserID uint, Record private.Calendar) error {
	err := global.EVA_DB.WithContext(ctx).Debug().Table("calendars").Where("id = ? and user_id = ?", Record.ID, UserID).Updates(map[string]interface{}{"is_all_day": Record.IsAllDay, "record_type": Record.RecordType, "start_time": Record.StartTime, "end_time": Record.EndTime, "title": Record.Title, "content": Record.Content, "alert": Record.Alert, "pre_alert_time": Record.PreAlertTime, "address": Record.Address, "user_status": Record.UserStatus, "repeat": Record.Repeat, "place": Record.Place, "img": Record.Img}).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *CalendarStore) UpdateEndTime(ctx context.Context, UserID uint, Record private.Calendar) error {
	err := global.EVA_DB.WithContext(ctx).Debug().Table("calendars").Where("id = ? and user_id = ?", Record.ID, UserID).Updates(map[string]interface{}{"end_time": Record.EndTime, "start_time": Record.StartTime}).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *CalendarStore) Detail(ctx context.Context, UserID uint, RecordID int) (private.Calendar, error) {
	result := private.Calendar{}
	if err := global.EVA_DB.WithContext(ctx).Where("id = ? and user_id = ?", RecordID, UserID).Find(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
