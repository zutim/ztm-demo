package app

import (
	logutil "github.com/zutim/log"
	"go.uber.org/zap"
)

func (a *apps) Logs(key string) *zap.SugaredLogger {
	if _, ok := a.Log[key]; ok {
		return a.Log[key]
	}
	return a.Log["log"]
}

func Log() *zap.SugaredLogger {
	return app.Logs("")
}

func newDefaultLogger() *zap.SugaredLogger {
	logger := logutil.InitLogger(func(options *logutil.LoggerOptions) {
		options.Path = "log/common.log"
	})
	defer logger.Sync()
	return logger.Sugar()
}

func OrderLog() *zap.SugaredLogger {
	return app.Logs("order")
}

func newOrderLogger() *zap.SugaredLogger {
	logger := logutil.InitLogger(func(options *logutil.LoggerOptions) {
		options.Path = "log/order.log"
	})
	defer logger.Sync()
	return logger.Sugar()
}

func UserLog() *zap.SugaredLogger {
	return app.Logs("user")
}

func newUserLogger() *zap.SugaredLogger {
	logger := logutil.InitLogger(func(options *logutil.LoggerOptions) {
		options.Path = "log/user.log"
	})
	defer logger.Sync()
	return logger.Sugar()
}
