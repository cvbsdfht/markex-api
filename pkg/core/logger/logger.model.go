package logger

import (
	"go.uber.org/zap"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})

	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Logger() *zap.Logger
}

type ZapLogger struct {
	ZapLogger *zap.Logger
}

type Options struct {
	FilePath      string
	Level         int
	Format        string
	ProdMode      bool
	IsDisplayTime bool
}