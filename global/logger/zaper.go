package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var log *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func init() {

	var writeSyncers []zapcore.WriteSyncer

	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	zapConfig := viper.Sub("zap")

	level := getLoggerLevel(viper.GetString(zapConfig.GetString("level")))

	// lumberjack 日志切割
	fileConfig := &lumberjack.Logger{
		Filename:   zapConfig.GetString("path"),    // 日志文件名
		MaxSize:    zapConfig.GetInt("maxSize"),    // 日志文件大小
		MaxAge:     zapConfig.GetInt("maxAge"),     // 最长保存天数
		MaxBackups: zapConfig.GetInt("maxBackups"), // 最多备份几个
		LocalTime:  zapConfig.GetBool("localTime"), // 日志时间戳
		Compress:   zapConfig.GetBool("compress"),  // 是否压缩文件，使用gzip
	}
	encoder := zap.NewProductionEncoderConfig()

	encoder.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000000"))
	}
	if zapConfig.GetBool("consoleStdout") {
		writeSyncers = append(writeSyncers, zapcore.AddSync(os.Stdout))
	}

	if zapConfig.GetBool("fileStdout") {
		writeSyncers = append(writeSyncers, zapcore.AddSync(fileConfig))
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder), zapcore.NewMultiWriteSyncer(writeSyncers...), zap.NewAtomicLevelAt(level))

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	log = logger.Sugar()
}

func getLoggerLevel(level string) zapcore.Level {
	if level, ok := levelMap[level]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func Debug(args ...interface{}) {
	log.Debug(args)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func DPanic(args ...interface{}) {
	log.DPanic(args...)
}

func DPanicf(format string, args ...interface{}) {
	log.DPanicf(format, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
