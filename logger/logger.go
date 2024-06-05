package logger

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/toolkits/file"
)

var logger *glog.Logger

func InitLog(name string) {

	logHome := filepath.Join(file.SelfDir(), "log")
	fileName := fmt.Sprintf("%s.{Y-m-d}.log", name)

	_logger := glog.New()
	_logger.SetConfigWithMap(g.Map{
		"path":                logHome,
		"level":               "all",
		"stdout":              true,
		"StStatus":            0,
		"stdoutColorDisabled": false,
		"file":                fileName,
	})

	_logger.SetStdoutPrint(true)
	_logger.SetFlags(44)

	// _logger.Info(context.Background(), "logger init success")
	logger = _logger

}

func SetLevel(level string) {
	switch level {
	case "debug":
		logger.SetLevel(glog.LEVEL_DEBU)
	case "info":
		logger.SetLevel(glog.LEVEL_INFO)
	case "warn":
		logger.SetLevel(glog.LEVEL_WARN)
	case "error":
		logger.SetLevel(glog.LEVEL_ERRO)
	case "fatal":
		logger.SetLevel(glog.LEVEL_FATA)
	default:
		logger.SetLevel(glog.LEVEL_ALL)
	}
}

func Debug(msg string) {
	logger.Debug(context.Background(), msg)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(context.Background(), format, args...)
}

func DebugWithCTX(ctx context.Context, msg string) {
	logger.Debug(ctx, msg)
}

func DebugfWithCTX(ctx context.Context, format string, args ...interface{}) {
	logger.Debugf(ctx, format, args...)
}

func Info(msg string) {
	logger.Info(context.Background(), msg)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(context.Background(), format, args...)
}

func InfoWithCTX(ctx context.Context, msg string) {
	logger.Info(ctx, msg)
}

func InfofWithCTX(ctx context.Context, format string, args ...interface{}) {
	logger.Infof(ctx, format, args...)
}

func Warn(msg string) {
	logger.Warning(context.Background(), msg)
}

func Warnf(format string, args ...interface{}) {
	logger.Warningf(context.Background(), format, args...)
}

func WarnWithCTX(ctx context.Context, msg string) {
	logger.Warning(ctx, msg)
}

func WarnfWithCTX(ctx context.Context, format string, args ...interface{}) {
	logger.Warningf(ctx, format, args...)
}

func Error(msg string) {
	logger.Error(context.Background(), msg)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(context.Background(), format, args...)
}

func ErrorWithCTX(ctx context.Context, msg string) {
	logger.Error(ctx, msg)
}

func ErrorfWithCTX(ctx context.Context, format string, args ...interface{}) {
	logger.Errorf(ctx, format, args...)
}

func Fatal(msg string) {
	logger.Fatal(context.Background(), msg)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(context.Background(), format, args...)
}

func FatalWithCTX(ctx context.Context, msg string) {
	logger.Fatal(ctx, msg)
}

func FatalfWithCTX(ctx context.Context, format string, args ...interface{}) {
	logger.Fatalf(ctx, format, args...)
}
