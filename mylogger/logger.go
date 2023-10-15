package mylogger

import (
	"context"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func DebugEnable() bool {
	return logger.level >= zapcore.DebugLevel
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zap.Field{Key: RequestIdKey, Type: zapcore.StringType, String: GetRequestID(ctx)})
	logger.Debug(msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zap.Field{Key: RequestIdKey, Type: zapcore.StringType, String: GetRequestID(ctx)})
	logger.Info(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zap.Field{Key: RequestIdKey, Type: zapcore.StringType, String: GetRequestID(ctx)})
	logger.Warn(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zap.Field{Key: RequestIdKey, Type: zapcore.StringType, String: GetRequestID(ctx)})
	logger.Error(msg, fields...)
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zap.Field{Key: RequestIdKey, Type: zapcore.StringType, String: GetRequestID(ctx)})
	logger.Fatal(msg, fields...)
}

func Debugf(ctx context.Context, template string, args ...interface{}) {
	template = strings.Join([]string{template, "|", RequestIdKey, ":", GetRequestID(ctx)}, "")
	logger.Debugf(template, args...)
}

func Infof(ctx context.Context, template string, args ...interface{}) {
	template = strings.Join([]string{template, "|", RequestIdKey, ":", GetRequestID(ctx)}, "")
	logger.Infof(template, args...)
}

func Warnf(ctx context.Context, template string, args ...interface{}) {
	template = strings.Join([]string{template, "|", RequestIdKey, ":", GetRequestID(ctx)}, "")
	logger.Warnf(template, args...)
}

func Errorf(ctx context.Context, template string, args ...interface{}) {
	template = strings.Join([]string{template, "|", RequestIdKey, ":", GetRequestID(ctx)}, "")
	logger.Errorf(template, args...)
}

func Fatalf(ctx context.Context, template string, args ...interface{}) {
	template = strings.Join([]string{template, "|", RequestIdKey, ":", GetRequestID(ctx)}, "")
	logger.Fatalf(template, args...)
}
