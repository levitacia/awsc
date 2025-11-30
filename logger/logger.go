package logger

import (
	"log"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var GlobalLogger *zap.Logger
var once sync.Once

func InitLogger(env string) {
	once.Do(func() {
		var config zap.Config

		if env == "production" {
			config = zap.NewProductionConfig()
			config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
			config.Level.SetLevel(zapcore.InfoLevel)
		} else {
			config = zap.NewDevelopmentConfig()
			config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

			config.Level.SetLevel(zapcore.DebugLevel)
		}

		var err error
		GlobalLogger, err = config.Build()
		if err != nil {
			log.Fatalf("Ошибка при инициализации Zap логгера: %v", err)
		}

		zap.ReplaceGlobals(GlobalLogger)
		GlobalLogger.Info("Zap logger инициализирован", zap.String("env", env))
	})
}

func L() *zap.Logger {
	if GlobalLogger == nil {
		l, _ := zap.NewDevelopment()
		return l
	}
	return GlobalLogger
}

func Sync() {
	if GlobalLogger != nil {
		if err := GlobalLogger.Sync(); err != nil && err.Error() != os.DevNull {
			log.Printf("Ошибка при очистке буфера логгера: %v", err)
		}
	}
}
