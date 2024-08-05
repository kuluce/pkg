package logger

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/kuluce/pkg/file"
)

var logger *glog.Logger

func InitLog(name string) {

	logHome := filepath.Join(file.SelfDir(), "log")
	fileName := fmt.Sprintf("%s.{Y-m-d}.log", name)

	_logger := glog.New()
	_logger.SetConfigWithMap(g.Map{
		"path":                 logHome,
		"level":                "all",
		"stdout":               true,
		"StStatus":             0,
		"stdoutColorDisabled":  false,
		"file":                 fileName,
		"rotateSize":           "100M",
		"rotateBackupLimit":    2,
		"rotateBackupExpire":   "7d",
		"rotateBackupCompress": 9,
	})

	_logger.SetStdoutPrint(true)
	_logger.SetFlags(glog.F_TIME_STD | glog.F_FILE_LONG)

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
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Debug(context.Background(), fmt.Sprintf("%s:%d: %s", file, line, msg))
}

func Debugf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Debugf(context.Background(), fmt.Sprintf("%s:%d: %s", file, line, format), args...)
}

func DebugWithCTX(ctx context.Context, msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Debug(ctx, fmt.Sprintf("%s:%d: %s", file, line, msg))
}

func DebugfWithCTX(ctx context.Context, format string, args ...interface{}) {

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Debugf(ctx, fmt.Sprintf("%s:%d: %s", file, line, format), args...)
}

func Info(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Info(context.Background(), fmt.Sprintf("%s:%d: %s", file, line, msg))
}

func Infof(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Infof(context.Background(), fmt.Sprintf("%s:%d: %s", file, line, format), args...)
}

func InfoWithCTX(ctx context.Context, msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Info(ctx, fmt.Sprintf("%s:%d: %s", file, line, msg))
}

func InfofWithCTX(ctx context.Context, format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Infof(ctx, fmt.Sprintf("%s:%d: %s", file, line, format), args...)
}

func Warn(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Warning(context.Background(), fmt.Sprintf("%s:%d: %s", file, line, msg))
}

func Warnf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Warningf(context.Background(), fmt.Sprintf("%s:%d: %s", file, line, format), args...)
}

func WarnWithCTX(ctx context.Context, msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Warning(ctx, fmt.Sprintf("%s:%d: %s", file, line, msg))
}

func WarnfWithCTX(ctx context.Context, format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Warningf(ctx, fmt.Sprintf("%s:%d: %s", file, line, format), args...)
}

func Error(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Error(context.Background(), fmt.Sprintf("%s:%d: %s", file, line, msg))
}

func Errorf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Errorf(context.Background(), fmt.Sprintf("%s:%d: %s", file, line, format), args...)
}

func ErrorWithCTX(ctx context.Context, msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Error(ctx, fmt.Sprintf("%s:%d: %s", file, line, msg))
}

func ErrorfWithCTX(ctx context.Context, format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Errorf(ctx, fmt.Sprintf("%s:%d: %s", file, line, format), args...)
}

func Fatal(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Fatal(context.Background(), fmt.Sprintf("%s:%d: %s", file, line, msg))
}

func Fatalf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Fatalf(context.Background(), fmt.Sprintf("%s:%d: %s", file, line, format), args...)
}

func FatalWithCTX(ctx context.Context, msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Fatal(ctx, fmt.Sprintf("%s:%d: %s", file, line, msg))
}

func FatalfWithCTX(ctx context.Context, format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, file = filepath.Split(file)

	logger.Fatalf(ctx, fmt.Sprintf("%s:%d: %s", file, line, format), args...)
}
