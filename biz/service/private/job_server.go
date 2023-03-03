package private

import (
	"eva/biz/service/send_message"
	"eva/biz/utils/timer"
	"eva/global"
	"eva/model/private"
	"fmt"
	"github.com/robfig/cron/v3"
	"strings"
	"time"
)

// 开启一个定时器，每小时去获取一次下一个小时内需要推送的任务。

var isFirst bool = true
var sendMessageService = send_message.SendMessageService{}

// 还需考虑，一个小时中如果有新增任务，该如何加入队列，可以nsq或kafka来做
// 也可以判断这个新增任务的时间是否属于当前小时内的，那么我们在入库的同时丢给刚才的 HandleJob 。
type JobService struct{}

func (j *JobService) SchedulerCalendarInit() {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()
	//var wg sync.WaitGroup
	for {
		if !isFirst {
			<-ticker.C
		}
		isFirst = false

		// 获取接下来一小时内需要发送的任务列表
		now := timer.GetCurrTime()
		h, _ := time.ParseDuration("1h")
		jobs, err := calendarJobService.GetJobsByTime(timer.TimeString(now), timer.TimeString(now.Add(1*h)))
		if err != nil {
			fmt.Printf("出错了：%v", err)
			return
		}
		// 任务通道
		ch := make(chan private.CalendarJob, 100)
		jobFunc := func(ch <-chan private.CalendarJob) {
			//wg.Add(100)
			for item := range ch {
				// 使用麻省理工的cron依赖
				// 也可以用带有延时队列功能的队列或数据库
				go j.HandleJob(item)
				//wg.Wait()
			}
		}
		// 处理任务
		go jobFunc(ch)
		// 投递任务
		for _, job := range jobs {
			ch <- job
		}
	}
}

func (j *JobService) HandleJob(job private.CalendarJob) {
	currentTime := job.NoticeTime.Format("05 04 15 2 1 2006")
	noticeTime := strings.Replace(currentTime, currentTime[len(currentTime)-4:len(currentTime)], "*", 1)
	fmt.Println("noticeTime:", noticeTime)
	cronID, _ := global.EVA_Timer.AddTaskByFunc("calendarSendAlert", noticeTime, func() {
		j.SendMessage(job)
	}, cron.WithSeconds())
	fmt.Println("cronID:", cronID)
	_ = calendarJobService.UpdateCronIDAndStatusById(int(job.ID), cronID, private.JobDoing)
}

func (j *JobService) SendMessage(job private.CalendarJob) {
	err := sendMessageService.SendDemo(job)
	// 是否成功
	isOk := private.JobSuccess
	if err != nil {
		isOk = private.JobFail
		fmt.Printf("通知失败:%v", err)
	}
	// 更新状态
	_ = calendarJobService.UpdateStatusById(int(job.ID), isOk)
	// 移除定时任务
	cronID, _ := calendarJobService.GetCronID(int(job.ID))
	global.EVA_Timer.Remove("CalendarSendAlert", int(cronID))
	//wg.Done()
}
