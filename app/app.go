package app

import (
	logutil "github.com/zutim/log"
	"github.com/zutim/pool"
	"go.uber.org/zap"
)

type Apps struct {
	Log  map[string]*zap.SugaredLogger
	Pool *pool.ScalableGroutinePool
}

func (a *Apps) Logs(key string) *zap.SugaredLogger {
	if _, ok := a.Log[key]; ok {
		return a.Log[key]
	}
	return a.Log["log"]
}

var App *Apps

func NewApps() *Apps {
	App = &Apps{
		Log: map[string]*zap.SugaredLogger{
			"log":   NewDefaultLogger(),
			"order": NewOrderLogger(),
			"user":  NewUserLogger(),
		},
		Pool: pool.NewGoroutinePool(),
	}
	return App
}

func NewDefaultLogger() *zap.SugaredLogger {
	logger := logutil.InitLogger(func(options *logutil.LoggerOptions) {
		options.Path = "log/common.log"
	})
	defer logger.Sync()
	return logger.Sugar()
}

func NewOrderLogger() *zap.SugaredLogger {
	logger := logutil.InitLogger(func(options *logutil.LoggerOptions) {
		options.Path = "log/order.log"
	})
	defer logger.Sync()
	return logger.Sugar()
}

func NewUserLogger() *zap.SugaredLogger {
	logger := logutil.InitLogger(func(options *logutil.LoggerOptions) {
		options.Path = "log/user.log"
	})
	defer logger.Sync()
	return logger.Sugar()
}
