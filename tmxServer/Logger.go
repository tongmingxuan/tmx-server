// Package tmxServer /*
package tmxServer

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
	"unsafe"
)

type Logger struct {
	ZapLogger *zap.Logger
}

func (l *Logger) Key() string {
	return "logger"
}

func (l *Logger) Handle() *BaseFrame {
	// 自定义zap日志配置
	encoderConfig := zap.NewProductionEncoderConfig()

	// 自定义时间格式
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05"))
	}
	// 创建普通的编译器
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	// 创建写入器
	writer := zapcore.AddSync(&lumberjack.Logger{
		//日志地址
		Filename: "runtime/logs/service.log",
		//日志大小
		MaxSize: 100,
		//保留日志个数
		MaxBackups: 5,
		//日志保留天数
		MaxAge: 2,
		//是否压缩
		Compress: false,
	})

	// new 创建的是zap的核心对象
	// 编译器，写入器，参数级别
	core := zapcore.NewCore(encoder, writer, zap.InfoLevel)

	l.ZapLogger = zap.New(core, zap.AddCaller(), zap.Hooks())

	return (*BaseFrame)(unsafe.Pointer(l))
}

// LoggerInfo
//
//	@Description: 默认日志记录
//	@receiver l
//	@param message 日志内容
//	@param tag 日志标识
//	@param data 日志相关参数
func (l *Logger) LoggerInfo(message string, tag string, data map[string]interface{}) {
	message = "tag:" + tag + ":message:" + message

	logData := []zap.Field{
		zap.String("tag", tag),
		zap.Any("message", message),
		zap.Any("log_time", time.Now().Format("2006-01-02 15:04:05")),
		zap.Any("data", data),
		zap.Any("trace", GetDebugTraceBySlice()),
	}

	l.ZapLogger.Info(message, logData...)
}
