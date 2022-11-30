package main

//
//import (
//	"fmt"
//	"time"
//
//	"code.byted.org/ad/polestar/biz/config"
//	"code.byted.org/gorm/bytedgen"
//	"code.byted.org/gorm/bytedgorm"
//	"gorm.io/gen"
//	"gorm.io/gorm"
//	"gorm.io/gorm/logger"
//)
//
//func main() {
//	config.InitAppConfig()
//	conf := &(config.AppConfig.DB)
//	// 与所写模块无关的请注释掉，避免entity标签修改！！！若需要生成或者更新model，只需要修改此行↓的数据库配置名和表名即可。提交代码时可以提交此文件的修改，也可以不提交。
//	// GenerateModel(&conf.AdQaETEData, "material_registration_monitor_fail_case")
//	// GenerateModel(&conf.AdQaETEData, "e2e_replay_task")
//	// GenerateModel(&conf.AdQaETEData, "e2e_realtime_task")
//	// GenerateModel(&conf.AdQaETEData, "e2e_data_task")
//	//GenerateModel(&conf.AdQaETEData, "e2e_data")
//	//GenerateModel(&conf.AdQaETEData, "e2e_auto_task")
//	GenerateModel(&conf.AdQaETEData, "e2e_auto_execution_result")
//	//GenerateModel(&conf.AdQaETEData, "e2e_auto_case")
//	//GenerateModel(&conf.AdQaETEData, "e2e_auto_business_case")
//	//GenerateModel(&conf.AdQaETEData, "e2e_auto_case_group")
//	//GenerateModel(&conf.AdQaETEData, "e2e_auto_test_plan")
//	//GenerateModel(&conf.AdQaETEData, "e2e_auto_business_run_detail")
//	// GenerateModel(&conf.AdQaETEData, "e2e_env")
//	// GenerateModel(&conf.AdQaETEData, "e2e_space_mgmt_info")
//	// GenerateModel(&conf.AdQaETEData, "monitor_e2e_tce_cluster")
//
//}
//
//func GenerateModel(conf *config.DatabaseConfig, tableName string) {
//	g := bytedgen.NewGenerator(gen.Config{
//		ModelPkgPath: "../../entity",
//	})
//
//	dbConf := conf
//	// reuse the database connection in Project or create a connection here
//	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary or it will panic
//	db, err := gorm.Open(
//		bytedgorm.MySQL(fmt.Sprintf("toutiao.mysql.%s", dbConf.DBName), dbConf.DBName).WithReadReplicas().With(func(conf *bytedgorm.DBConfig) {
//			// 使用自定义的 DSN
//			conf.DBHostname = fmt.Sprintf("toutiao.mysql.%s", dbConf.DBName) + "_read"
//			conf.ReadTimeout = 2 * time.Second
//		}),
//		bytedgorm.Logger{LogLevel: logger.Info, IgnoreRecordNotFoundError: true},
//		bytedgorm.WithDefaults(),
//	)
//	//db, err := gorm.Open(bytedgorm.MySQL("toutiao.mysql.ad_qa_quantum" /*数据库PSM*/, "ad_qa_quantum" /*数据库名*/).With(func(conf *bytedgorm.DBConfig) {
//	//	// 使用自定义的 DSN
//	//	conf.ReadTimeout = 5 * time.Second
//	//	conf.Timeout = 5 * time.Second
//	//	if dbConf.Type == "NORMAL" {
//	//		//"ad_audit_camp_w:0xGXXGCWEl4SF8c_NKPzMWoCxZQPdTU7@tcp(10.231.26.218:3306)/ad_audit_campanula?charset=utf8mb4&parseTime=True&loc=Local"
//	//		dsnTemplate := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
//	//		conf.DSN = fmt.Sprintf(dsnTemplate, dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DBName)
//	//		//logs.Info("using normal dsn to create db connetion, dsn = %s", conf.DSN)
//	//	}
//	//
//	//}), bytedgorm.WithDefaults(), bytedgorm.WithSingularTable())
//	////RegisterTables(DB)
//
//	if err != nil {
//		panic(err)
//	}
//
//	g.UseDB(db)
//	g.GenerateModel(tableName)
//
//	g.Execute()
//}
