package logg

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogger() {
	var err error
	Logger = Config()

	if err != nil {
		panic("Ошибка инициализации логгера: " + err.Error())
	}
}
