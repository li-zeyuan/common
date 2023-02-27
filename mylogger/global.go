package mylogger

import "go.uber.org/zap"

var logger *Logger

func Init(cfg *LoggerCfg) error {
	options := make([]optionFun, 0)
	if cfg == nil {
		options = defaultOptions
	} else {
		options = append(options,
			WhitLevel(cfg.Level),
			WhitLoggingDir(cfg.LoggingDir),
			WhitMaxSize(cfg.MaxSize),
			WhitMaxAge(cfg.MaxAge),
			WhitMaxBackup(cfg.MaxBackup),
			WhitIsCompress(cfg.IsCompress),
			WhitIsConsole(cfg.IsConsole),
		)
	}

	for _, o := range options {
		o(cfg)
	}

	var err error
	logger, err = NewZap()
	return err
}

func GetZapLogger() *zap.Logger {
	return logger.log
}
