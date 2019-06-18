package logging

import (
	"github.com/dncmn/xorm-log"
)

var (
	logging *xorm_log.DiscardLogger
)

func init() {
	xorm_log.Init(xorm_log.FileLogConfig{
		Path:     "/var/log/game-server",
		Filename: "/var/log/game-server/app.log",
		MaxLines: 1000000,
		Maxsize:  1 << 28, //256 MB
		Daily:    true,
		MaxDays:  3,
		Rotate:   true,
		Level:    xorm_log.DebugLevel,
	})

	logging = xorm_log.GetLogger()
	logging.Info("logger success init.")
}

func GetLogger() *xorm_log.DiscardLogger {
	return logging
}
