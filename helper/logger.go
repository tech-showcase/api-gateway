package helper

import (
	"github.com/go-kit/kit/log"
	"os"
)

var LoggerInstance log.Logger

func NewLogger() log.Logger {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	return logger
}
