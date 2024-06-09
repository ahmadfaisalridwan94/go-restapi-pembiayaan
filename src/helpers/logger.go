package helpers

import (
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
)

func CreateAppLog(data interface{}) *logrus.Entry {
	l := logrus.New()
	// hook, err := lSyslog.NewSyslogHook("tcp", GetEnv("SYSLOG_PATH", ""), syslog.LOG_INFO, "user-service")

	// if err == nil {
	// 	l.Out = io.Discard
	// 	l.Hooks.Add(hook)
	// }

	l.SetReportCaller(true)

	addData := logrus.Fields{}
	mapstructure.Decode(data, &addData)

	logData := logrus.Fields{}
	logData["interface"] = GetEnv("INTERFACE", "")
	logData["data"] = addData

	return l.WithFields(logData)
}
