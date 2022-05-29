package log

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

const logTmFmt = "2006-01-02 15:04:05.000"

// 创建日志对象
func GetLogger() *zap.Logger {
	core := zapcore.NewTee(
		zapcore.NewCore(getEncoder(), zapcore.Lock(os.Stdout), zapcore.DebugLevel), // 写入控制台
		//zapcore.NewCore(getEncoder(), getWriteSyncer(), getLevelEnabler()),                // 写入文件
	)
	caller := zap.AddCaller()
	development := zap.Development()
	return zap.New(core, caller, development)
}

// 日志输出格式
func getEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(
		zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level_name",
			TimeKey:        "ts",
			NameKey:        "logger",
			CallerKey:      "caller_line", // 打印文件名和行数
			FunctionKey:    zapcore.OmitKey,
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,      // 默认换行符"\n"
			EncodeLevel:    cEncodeLevel,                   // 日志等级
			EncodeTime:     cEncodeTime,                    // 日志时间格式显示
			EncodeDuration: zapcore.SecondsDurationEncoder, // 时间序列化，Duration为经过的浮点秒数
			EncodeCaller:   cEncodeCaller,                  // 日志行号显示
			EncodeName:     zapcore.FullNameEncoder,
		})
}

// 日志输出到文件中
func getWriteSyncer() zapcore.WriteSyncer {
	today := time.Now().Format("2006-01-02")
	file, _ := os.Create(fmt.Sprintf("log/%s.log", today))
	return zapcore.AddSync(file)
}

// 输出日志级别
func getLevelEnabler() zapcore.Level {
	level := viper.GetString("log.level")
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel // 只会打印出info及其以上级别的日志
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

// cEncodeLevel 自定义日志级别显示
func cEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

// cEncodeTime 自定义时间格式显示
func cEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format(logTmFmt) + "]")
}

// cEncodeCaller 自定义行号显示
func cEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}
