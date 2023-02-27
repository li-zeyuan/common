package mylogger

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

func getWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    config.MaxSize, // Mb
		MaxBackups: config.MaxBackup,
		MaxAge:     config.MaxAge,
		Compress:   config.IsCompress,
	}
}

