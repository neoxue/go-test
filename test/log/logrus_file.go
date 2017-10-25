package main

import (
	"github.com/Sirupsen/logrus"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
)

func main() {
		//rl, _ := rotatelogs.New("/data1/ms/log/access_log.%Y%m%d%H%M")
		rl, _ := rotatelogs.New("/data1/ms/log/access_log.%Y%m%d")
		logrus.SetOutput(rl)

		/* elsewhere ... */
		logrus.Printf("Hello, World!")

}

var Log *logrus.Logger

/*
func NewLogger( config map[string]interface{} ) *logrus.Logger {
	if Log != nil {
		return Log
	}

	Log = logrus.New()
	Log.Formatter = &logrus.JSONFormatter{}

	path := "/var/log/go.log"
	writer := rotatelogs.New(
		path + ".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(86400) * time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800) * time.Second),
	)

	log.Out = writer


	var infoLevel logrus.Level
	infoLevel = logrus.DebugLevel
	a := lfshook.PathMap{
		infoLevel: "/data1/ms/log/logrus_info.log",
	}
	hook := lfshook.NewHook(a)
	Log.Hooks.Add(hook)
	return Log
}
*/