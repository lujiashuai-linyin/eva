package private

import (
	"eva/biz/service/send_message"
	"eva/biz/utils/timer"
	"eva/model/private"
	"fmt"
	"time"
)

// 开启一个定时器，每小时去获取一次下一个小时内需要推送的任务。

var isFirst bool = true
var sendMessageService = send_message.SendMessageService{}

// 还需考虑，一个小时中如果有新增任务，该如何加入队列，可以nsq或kafka来做
// 也可以这个新增任务的时间属于当前小时内的，那么我们在入库的同时丢给刚才的 HandleJob 。
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
				// 发送通知, 切片处理100，前提是任务列表要按照时间顺序排列
				//go HandleJob(item, &wg)
				// 量级不大，就简单点
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
	now := timer.GetCurrTime()
	diff := job.NoticeTime.Sub(now)
	//fmt.Println("diff:", diff)
	timerJob := time.NewTimer(diff)
	_ = calendarJobService.UpdateStatusById(int(job.ID), private.JobDoing)
	<-timerJob.C

	err := sendMessageService.SendDemo(job)
	//sendTool = &sendMessageService.SmsMsg{Job: job}
	//if job.Phone == "" {
	//	sendTool = &sendMessageService.EmailMsg{Job: job}
	//}
	//err := sendMessageService.Notice(sendTool)
	//成功与否
	isOk := private.JobSuccess
	if err != nil {
		isOk = private.JobFail
		fmt.Printf("通知失败:%v", err)
	}
	_ = calendarJobService.UpdateStatusById(int(job.ID), isOk)
	//wg.Done()
}
