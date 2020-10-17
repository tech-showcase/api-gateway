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

func NewFileLogger(filepath string) (log.Logger, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}

	var logger log.Logger
	logger = log.NewLogfmtLogger(f)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	return logger, nil
}
