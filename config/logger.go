package config

import (
	"bytes"
	"fmt"
	"gin-base/common/extend/context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Log 日志
type Log struct {
	Access     bool `yaml:"access"`      // 是否记录访问日志
	MaxSize    int  `yaml:"max-size"`    // 单个日志文件大小（MB）
	MaxBackups int  `yaml:"max-backups"` // 最多保留的旧日志文件数
	MaxDay     int  `yaml:"max-day"`     // 保留的最大天数
}

// Logger 包装器
type Logger struct {
	*zap.Logger
}

// NewLogger 创建新的 Logger 包装器
func NewLogger(zapLogger *zap.Logger) *Logger {
	return &Logger{zapLogger}
}

// GetLogFields 记录字段
func GetLogFields(fields []zap.Field) []zap.Field {
	c := context.GetGinContext("GinLogger")
	if c == nil {
		return nil
	}

	ip := c.ClientIP()
	path := c.Request.URL.Path
	method := c.Request.Method
	var params string

	switch method {
	case http.MethodGet, http.MethodDelete:
		params = c.Request.URL.RawQuery
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		// 读取 Body
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err == nil {
			params = string(bodyBytes)
			// 恢复 Body 的原始状态
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}
	}

	// 清理参数中的空格和换行符
	params = strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(params, "\r", ""), "\n", ""))

	// 添加日志字段
	fields = append(fields,
		zap.String("ip", ip),
		zap.String("method", method),
		zap.String("path", path),
		zap.String("params", params),
		zap.Object("stacktrace", StackTrace{}),
	)

	return fields
}

type StackTrace struct{}

// MarshalLogObject 格式化堆栈信息
func (s StackTrace) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	// 获取堆栈信息
	stack := zap.Stack("stack").String
	lines := strings.Split(stack, "\n")

	_ = encoder.AddArray("stack", zapcore.ArrayMarshalerFunc(func(arrEnc zapcore.ArrayEncoder) error {
		for _, line := range lines {
			arrEnc.AppendString(strings.TrimSpace(line))
		}
		return nil
	}))

	return nil
}

// addContextFields 添加上下文字段
func (l *Logger) addContextFields(fields []zap.Field) []zap.Field {
	for _, field := range fields {
		if field.Key == "ip" {
			return fields
		}
	}

	return GetLogFields(fields)
}

// Debug 包装器方法
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	fields = l.addContextFields(fields)
	l.Logger.Debug(msg, fields...)
}

// Info 包装器方法
func (l *Logger) Info(msg string, fields ...zap.Field) {
	fields = l.addContextFields(fields)
	l.Logger.Info(msg, fields...)
}

// Warn 包装器方法
func (l *Logger) Warn(msg string, fields ...zap.Field) {
	fields = l.addContextFields(fields)
	l.Logger.Warn(msg, fields...)
}

// Error 包装器方法
func (l *Logger) Error(msg string, fields ...zap.Field) {
	fields = l.addContextFields(fields)
	l.Logger.Error(msg, fields...)
}

// DPanic 包装器方法
func (l *Logger) DPanic(msg string, fields ...zap.Field) {
	fields = l.addContextFields(fields)
	l.Logger.DPanic(msg, fields...)
}

// Panic 包装器方法
func (l *Logger) Panic(msg string, fields ...zap.Field) {
	fields = l.addContextFields(fields)
	l.Logger.Panic(msg, fields...)
}

// Fatal 包装器方法
func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	fields = l.addContextFields(fields)
	l.Logger.Fatal(msg, fields...)
}

// InitLogger 初始化日志
func InitLogger() *Logger {
	// 设置日志文件输出路径和文件年月日
	logPath := "log/" + time.Now().Format("2006-01-02") + ".log"

	// 配置 Lumberjack 日志轮转
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    config.Log.MaxSize,    // 单个日志文件大小（MB）
		MaxBackups: config.Log.MaxBackups, // 最多保留的旧日志文件数
		MaxAge:     config.Log.MaxDay,     // 保留的最大天数
		Compress:   true,
	}
	fmt.Printf("日志配置：%v\n", config.Log)

	// 创建一个Zap配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(lumberJackLogger),
		atomicLevel,
	)

	// 添加 Caller（文件和行号）
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return NewLogger(zapLogger)
}
