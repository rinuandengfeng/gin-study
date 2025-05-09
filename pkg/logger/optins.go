package logger

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	// DefaultLevel 默认日志级别.
	DefaultLevel = zapcore.InfoLevel
	// DefaultTimeLayout 默认时间格式.
	DefaultTimeLayout = time.DateTime
)

type (
	// Option 选项卡配置方法, 定义日志启动时需要使用的功能.
	Option func(*options)
	// options 日志相关配置信息.
	options struct {
		file              io.Writer
		ioClose           io.Closer
		timeLayout        string
		traceID           string
		ctxFields         []string
		level             zapcore.Level
		disableConsole    bool
		disableCaller     bool
		disableStacktrace bool
	}
)

func newDefaultOptions() *options {
	return &options{
		level:             DefaultLevel,
		file:              nil,
		timeLayout:        DefaultTimeLayout,
		disableConsole:    false,
		disableCaller:     false,
		disableStacktrace: false,
		ctxFields:         make([]string, 0),
		traceID:           "",
	}
}

// WithDebugLevel 设置日志级别为 Debug.
func WithDebugLevel() Option {
	return func(opt *options) {
		opt.level = zapcore.DebugLevel
	}
}

// WithInfoLevel 设置日志级别为Info.
func WithInfoLevel() Option {
	return func(opt *options) {
		opt.level = zapcore.InfoLevel
	}
}

// WithWarnLevel 设置日志级别为Warn.
func WithWarnLevel() Option {
	return func(opt *options) {
		opt.level = zapcore.WarnLevel
	}
}

// WithErrorLevel 设置日志级别为Error.
func WithErrorLevel() Option {
	return func(opt *options) {
		opt.level = zapcore.ErrorLevel
	}
}

// WithPanicLevel 设置日志级别为Panic.
func WithPanicLevel() Option {
	return func(opt *options) {
		opt.level = zapcore.PanicLevel
	}
}

// WithFatalLevel 设置日志级别为Fatal.
func WithFatalLevel() Option {
	return func(opt *options) {
		opt.level = zapcore.FatalLevel
	}
}

// WithTimeLayout 设置日志打印时间格式.
func WithTimeLayout(layout string) Option {
	return func(opt *options) {
		opt.timeLayout = layout
	}
}

// WithFilePath 通过文件路径进行日志写入.
func WithFilePath(file string) Option {
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, 0o766); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0o766)
	if err != nil {
		panic(err)
	}

	return func(opt *options) {
		opt.file = zapcore.Lock(f)
	}
}

// WithFileP 自定义写入配置文件.
func WithFileP(file io.Writer) Option {
	return func(opt *options) {
		opt.file = file
	}
}

// 使用 lumberjack.Logger 做切割配置.
func WithFileRotationP(file string) Option {
	return func(opt *options) {
		// lumberjack 需要额外返回一个close,确保进程推出前将写入的内容刷到磁盘
		lumbLog := &lumberjack.Logger{
			Filename:   file, // 文件路径
			MaxSize:    128,  // 单个文件最大尺寸， 默认单位 MB
			MaxAge:     30,   // 最大保存时间， 默认单位天
			MaxBackups: 300,  // 最多保留300个备份
			LocalTime:  true, // 使用本地之间
			Compress:   true, // 是否压缩, 默认不压缩
		}

		opt.file = lumbLog
		opt.ioClose = lumbLog
	}
}

// WithDisableConsole 禁止写入控制台.
func WithDisableConsole() Option {
	return func(opt *options) {
		opt.disableConsole = true
	}
}

// WithDisableCaller 禁止在日志中显示调用日志所在的文件和行号.
func WithDisableCaller() Option {
	return func(opt *options) {
		opt.disableCaller = true
	}
}

// WithDisableStacktrace 禁止在panic及以上级别答应堆栈信息.
func WithDisableStacktrace() Option {
	return func(opt *options) {
		opt.disableStacktrace = true
	}
}

// WithCtxFields 设置从上下文获取的字段.
func WithCtxFields(fields []string) Option {
	return func(opt *options) {
		opt.ctxFields = fields
	}
}

// WithTraceID 日志追踪字段，需要在上下文中进行传递.
func WithTraceID(traceID string) Option {
	return func(opt *options) {
		opt.traceID = traceID
	}
}
