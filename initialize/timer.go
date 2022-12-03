package initialize

import (
	"eva/biz/config"
	"eva/biz/utils"
	"eva/global"
	"fmt"
	"github.com/robfig/cron/v3"
)

func Timer() {
	if global.EVA_CONFIG.Timer.Start {
		for i := range global.EVA_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.EVA_CONFIG.Timer.WithSeconds {
					option = append(option, cron.WithSeconds())
				}
				_, err := global.EVA_Timer.AddTaskByFunc("ClearDB", global.EVA_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.EVA_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				}, option...)
				if err != nil {
					fmt.Println("add timer error:", err)
				}
			}(global.EVA_CONFIG.Timer.Detail[i])
		}
	}
}
