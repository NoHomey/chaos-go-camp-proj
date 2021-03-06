package logcrtr

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

//Config configures new zap.Logger.
func Config(path string) *zap.Logger {
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(
			&lumberjack.Logger{
				Filename:   path,
				MaxSize:    maxSizeInMB,
				MaxBackups: maxBackups,
				MaxAge:     maxAgeInDays,
			},
		),
		zapcore.InfoLevel,
	)
	return zap.New(core)
}

const (
	maxSizeInMB  = 500
	maxBackups   = 3
	maxAgeInDays = 14
)
