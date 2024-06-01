package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// https://nyogjtrc.github.io/posts/2019/09/log-rotate-with-zap-logger/
func NewLogger(option *Options) Logger {
	var ioWriter = &lumberjack.Logger{
		Filename:   option.FilePath,
		MaxSize:    10,
		MaxBackups: 90,
		MaxAge:     90,
		LocalTime:  true,
		Compress:   true,
	}

	writeFile := zapcore.AddSync(ioWriter)
	writeStdout := zapcore.AddSync(os.Stdout)

	if option.ProdMode {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
		if !option.IsDisplayTime {
			encoderConfig.TimeKey = "" // remove time
		}
		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(writeFile, writeStdout),
			zap.InfoLevel,
		)

		return &ZapLogger{
			ZapLogger: zap.New(
				core,
				zap.AddCaller(),
				zap.AddCallerSkip(1),
				zap.AddStacktrace(zap.ErrorLevel),
			),
		}
	} else {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
		if !option.IsDisplayTime {
			encoderConfig.TimeKey = "" // remove time
		}
		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(writeFile, writeStdout),
			zap.DebugLevel,
		)
		
		return &ZapLogger{
			ZapLogger: zap.New(
				core,
				zap.AddCaller(),
				zap.AddCallerSkip(1),
				zap.AddStacktrace(zap.ErrorLevel),
			),
		}
	}
}

// Debug implements Logger.
func (z *ZapLogger) Debug(args ...interface{}) {
	z.ZapLogger.Sugar().Debug(args)
}

// Error implements Logger.
func (z *ZapLogger) Error(args ...interface{}) {
	z.ZapLogger.Sugar().Error(args)
}

// Info implements Logger.
func (z *ZapLogger) Info(args ...interface{}) {
	z.ZapLogger.Sugar().Info(args)
}

// Warn implements Logger.
func (z *ZapLogger) Warn(args ...interface{}) {
	z.ZapLogger.Sugar().Warn(args)
}

// Debugf implements Logger.
func (z *ZapLogger) Debugf(template string, args ...interface{}) {
	z.ZapLogger.Sugar().Debugf(template, args...)
}

// Errorf implements Logger.
func (z *ZapLogger) Errorf(template string, args ...interface{}) {
	z.ZapLogger.Sugar().Errorf(template, args...)
}

// Infof implements Logger.
func (z *ZapLogger) Infof(template string, args ...interface{}) {
	z.ZapLogger.Sugar().Infof(template, args...)
}

// Warnf implements Logger.
func (z *ZapLogger) Warnf(template string, args ...interface{}) {
	z.ZapLogger.Sugar().Warnf(template, args...)
}

// Logger implements Logger.
func (z *ZapLogger) Logger() *zap.Logger {
	return z.ZapLogger
}
