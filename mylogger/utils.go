package mylogger

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

func getWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    option.MaxSize, // Mb
		MaxBackups: option.MaxBackup,
		MaxAge:     option.MaxAge,
		Compress:   option.IsCompress,
	}
}

