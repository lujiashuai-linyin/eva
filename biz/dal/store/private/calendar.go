package private

import (
	"context"
	"eva/global"
	"eva/model/private"
)

type CalendarStoreIf interface {
	FindRecordList(ctx context.Context, UserID int) ([]*private.Calendar, error)
	DeleteRecord(ctx context.Context, RecordID int) error
	Create(ctx context.Context, Record private.Calendar) error
	Update(ctx context.Context, Record private.Calendar) error
}

type CalendarStore struct{}

// 创建一条事件

func (store *CalendarStore) Create(ctx context.Context, Record private.Calendar) error {
	//err := global.EVA_DB.Transaction(func(tx *gorm.DB) error {
	//	for _, record := range Record {
	//		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
	//		if err := tx.Create(&record).Error; err != nil {
	//			// 返回任何错误都会回滚事务
	//			return err
	//		}
	//	}
	//	return nil
	//})
	if err := global.EVA_DB.Create(&Record).Error; err != nil {
		return err
	}
	return nil
}

// 获取事件列表

func (store *CalendarStore) FindRecordList(ctx context.Context, UserID int) ([]*private.Calendar, error) {
	result := []*private.Calendar{}
	if err := global.EVA_DB.WithContext(ctx).Where("user_id = ?", UserID).Find(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
func (store *CalendarStore) DeleteRecord(ctx context.Context, RecordID int) error {
	if err := global.EVA_DB.WithContext(ctx).Where("id = ?", RecordID).Delete(&private.Calendar{}).Error; err != nil {
		return err
	}
	return nil
}
func (store *CalendarStore) Update(ctx context.Context, Record private.Calendar) error {
	err := global.EVA_DB.WithContext(ctx).Debug().Where("id = ?", Record.ID).Updates(map[string]interface{}{"is_all_day": Record.IsAllDay, "record_type": Record.RecordType, "start_time": Record.StartTime, "end_time": Record.EndTime, "title": Record.Title, "content": Record.Content, "alert": Record.Alert, "pre_alert_time": Record.PreAlertTime, "address": Record.Address, "user_status": Record.UserStatus, "repeat": Record.Repeat, "place": Record.Place, "img": Record.Img}).Error
	if err != nil {
		return err
	}
	return nil
}
