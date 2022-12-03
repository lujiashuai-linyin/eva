//// Code generated by hertztool.
package main

import (
	"eva/model/private"
	"fmt"
	"time"
)

//
//import (
//	"code.byted.org/ad/polestar/entity"
//	"code.byted.org/gorm/bytedgen"
//	"gorm.io/gen"
//)
//
//// GEN Guideline: https://bytedance.feishu.cn/wiki/wikcnbYLEL78aOYLBsT2ExUGiEd
//
//// generated query code
//func main() {
//
//	g := bytedgen.NewGenerator(gen.Config{
//		OutPath: "../../biz/dal/gen/",
//	})
//
//	g.ApplyBasic(entity.RuleList{}, entity.RuleGray{}, entity.ActionInfo{})
//
//	g.ApplyBasic(entity.MaterialRegistrationMonitorFailCase{})
//
//	g.ApplyBasic(entity.E2eEnv{})
//	g.ApplyBasic(entity.E2eReplayTask{})
//	g.ApplyBasic(entity.E2eRealtimeTask{})
//	g.ApplyBasic(entity.E2eDataTask{})
//	g.ApplyBasic(entity.E2eDatum{})
//	g.ApplyBasic(entity.E2eSpaceMgmtInfo{})
//	g.ApplyBasic(entity.MonitorE2eTceCluster{})
//	g.ApplyBasic(entity.E2eAutoTask{})
//	g.ApplyBasic(entity.E2eAutoExecutionResult{})
//	g.ApplyBasic(entity.E2eAutoCase{})
//	g.ApplyBasic(entity.E2eAutoCaseGroup{})
//	g.ApplyBasic(entity.E2eAutoCasePath{})
//	g.ApplyBasic(entity.E2eAutoTestPlan{})
//	g.ApplyBasic(entity.E2eAutoBusinessCase{})
//	g.ApplyBasic(entity.E2eAutoBusinessRunDetail{})
//
//	g.Execute()
//}

func main() {

	// Declaring Time in UTC
	hh, _ := time.ParseDuration("1h")
	record := private.Calendar{StartTime: time.Now(), EndTime: time.Now().Add(hh)}
	recordList := []*private.Calendar{}
	i := 0
	recordList = append(recordList, &record)
	for ; i < 6; i++ {
		record.StartTime.Add(4 * hh)
		record.EndTime.Add(4 * hh)
		record := record
		fmt.Println(record.EndTime)
		recordList = append(recordList, &record)
	}
	fmt.Println(recordList)
	for _, i := range recordList {
		fmt.Println(*i)
	}
}