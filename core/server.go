package core

import (
	"fmt"
	"time"

	"eva/global"
	"eva/initialize"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.EVA_CONFIG.System.UseMultipoint || global.EVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	//if global.EVA_DB != nil {
	//	system.LoadAll()
	//}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.EVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.EVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	我是eva！！我一直期待着与您再次相见。
	eva当前年龄: 1.0
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, address)
	global.EVA_LOG.Error(s.ListenAndServe().Error())
}
