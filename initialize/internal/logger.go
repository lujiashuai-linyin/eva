package internal

import (
	"fmt"

	"eva/global"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
// Author [SliverHorn](https://github.com/SliverHorn)
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
// Author [SliverHorn](https://github.com/SliverHorn)
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.EVA_CONFIG.System.DbType {
	case "mysql":
		logZap = global.EVA_CONFIG.Mysql.LogZap
	case "pgsql":
		logZap = global.EVA_CONFIG.Pgsql.LogZap
	}
	if logZap {
		global.EVA_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
