package main

import (
	"eva/core"
	"eva/global"
	"eva/initialize"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	global.EVA_VP = core.Viper() // 初始化Viper
	global.EVA_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.EVA_LOG)
	global.EVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.InitStore()
	//initialize.Timer()
	//initialize.DBList()
	if global.EVA_DB != nil {
		initialize.RegisterTables(global.EVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.EVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
