package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Module is one main modules
var Module = fx.Options(
	fx.Provide(New),
)

// Logger ...
type Logger struct {
	*zap.Logger
}

// New create new Logger type
func New() *Logger {

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "log/app.log",
		MaxSize:    10,
		MaxBackups: 10,
		MaxAge:     10,
	})

	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "datetime"
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		zap.DebugLevel,
	)

	return &Logger{zap.New(core)}
}

func (l *Logger) Log(message, service, function, operation string) {
	l.Info(message,
		zap.String("Service", service),
		zap.String("Function", function),
		zap.String("Operation", operation),
	)
}
