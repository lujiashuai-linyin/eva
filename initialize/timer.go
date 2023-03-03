package initialize

//func Timer() {
//	if global.EVA_CONFIG.Timer.Start {
//		currentTime := time.Now().Format("05 04 15 2 1 2006")
//		currentTime = strings.Replace(currentTime, currentTime[len(currentTime)-4:len(currentTime)], "*", 1)
//		fmt.Println(currentTime)
//		id, err := global.EVA_Timer.AddTaskByFunc("CalendarSendAlert", "20 48 22 3 12 *", func() {
//			fmt.Printf("发送消息定时")
//
//		}, cron.WithSeconds())
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Println("conID:", id)
//		global.EVA_Timer.Remove("CalendarSendAlert", int(id))
//		//for i := range global.EVA_CONFIG.Timer.Detail {
//		//go func(detail config.Detail) {
//		//	var option []cron.Option
//		//	if global.EVA_CONFIG.Timer.WithSeconds {
//		//		option = append(option, cron.WithSeconds())
//		//	}
//		//	_, err := global.EVA_Timer.AddTaskByFunc("ClearDB", global.EVA_CONFIG.Timer.Spec, func() {
//		//		err := utils.ClearTable(global.EVA_DB, detail.TableName, detail.CompareField, detail.Interval)
//		//		if err != nil {
//		//			fmt.Println("timer error:", err)
//		//		}
//		//	}, option...)
//		//	if err != nil {
//		//		fmt.Println("add timer error:", err)
//		//	}
//		//}(global.EVA_CONFIG.Timer.Detail[i])
//		//}
//	}
//}
