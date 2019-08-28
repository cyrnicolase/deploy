package logger

import (
	"deploy/config"

	clog "github.com/cyrnicolase/logger"
)

// Log 日志记录器
var Log *clog.FileLogger

// Boot 启动
func Boot() {
	logConfig := config.App.Log
	filename := logConfig.Path + "/" + logConfig.Name
	log, err := clog.NewFileLogger(filename, logConfig.Level)
	if nil != err {
		panic(err.Error())
	}

	Log = log
}
