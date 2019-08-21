package logger

import (
	"fmt"
	"io"
	"os"

	"deploy/config"

	"github.com/sirupsen/logrus"
)

// Log 日志
var Log = logrus.New()
var logOutput io.Writer

func init() {
	logOutput = os.Stdout
}

// Boot 启动
func Boot() {
	logOutput, err := fileOutput()
	if nil != err {
		fmt.Println("日志文件打开失败")
		return
	}

	// Log.SetFormatter(&logrus.TextFormatter{})
	Log.SetFormatter(&MixFormatter{Filename: config.App.Log.Name})
	Log.SetReportCaller(true)
	Log.SetOutput(logOutput)
	Log.SetLevel(logrus.Level(config.App.Log.Level))
}

// Info info
func Info(msg string, hash map[string]interface{}) {
	Log.WithFields(logrus.Fields(hash)).Info(msg)
}

// Warn warn
func Warn(msg string, hash map[string]interface{}) {
	Log.WithFields(logrus.Fields(hash)).Warn(msg)
}

// Error Error
func Error(msg string, hash map[string]interface{}) {
	Log.WithFields(logrus.Fields(hash)).Error(msg)
}

// Fatal Fatal
func Fatal(msg string, hash map[string]interface{}) {
	Log.WithFields(logrus.Fields(hash)).Fatal(msg)
}

func fileOutput() (io.Writer, error) {
	filename := generateOutputFilename()
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if nil != err {
		return nil, err
	}

	return file, nil
}

func generateOutputFilename() string {
	path := config.App.Log.Path
	name := config.App.Log.Name
	// mode := config.App.Log.Mode

	return fmt.Sprintf("%s/%s", path, name)
}
