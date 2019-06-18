package logging

import (
	"github.com/snaplingo-org/sllog"
)

var (
	logging *sllog.Sllogger
)

func init() {
	sllog.Init(sllog.FileLogConfig{
		Path:     "/var/log/golang/web-macaron",
		Filename: "/var/log/golang/web-macaron.log",
		MaxLines: 1000000,
		Maxsize:  1 << 28, //256 MB
		Daily:    true,
		MaxDays:  30,
		Rotate:   true,
		Level:    sllog.DebugLevel,
	})

	logging = sllog.GetLogger()
	logging.Info("logger success init.")
}
func GetLogger() *sllog.Sllogger {
	return logging
}
