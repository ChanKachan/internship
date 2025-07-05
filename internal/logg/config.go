package logg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Config() *zap.Logger {
	config := zap.NewProductionEncoderConfig()

	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.TimeKey = "time"
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.AddSync(os.Stdout), // вывод в консоль
		zapcore.InfoLevel,
	))

	return logger
}
