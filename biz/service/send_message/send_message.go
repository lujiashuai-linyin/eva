package send_message

import (
	"eva/biz/utils/timer"
	"eva/model/private"
	"fmt"
	"time"
)

type SendMessageService struct{}

func (s *SendMessageService) SendDemo(job private.CalendarJob) error {
	now := time.Now().Format(timer.TimeFormat)
	fmt.Println("发送成功:\n时间-内容", now, job.Content)
	return nil
}
