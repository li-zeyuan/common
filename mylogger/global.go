package mylogger

import "go.uber.org/zap"

var logger *Logger

func Init(options ...optionFun) error {
	if len(options) == 0 {
		options = defaultOptions
	}

	for _, o := range options{
		o(option)
	}

	var err error
	logger, err = NewZap()
	return err
}

func GetZapLogger() *zap.Logger {
	return logger.log
}