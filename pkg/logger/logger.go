package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
)

type (
	// ILogger 定义服务的日记接口, 该接口只包含了支持的日志记录方式.
	ILogger interface {
		Debug(context.Context, string, ...any)
		Info(context.Context, string, ...any)
		Warn(context.Context, string, ...any)
		Error(context.Context, string, ...any)
		Panic(context.Context, string, ...any)
		Fatal(context.Context, string, ...any)
		Sync()
	}

	ZapLogger struct {
		ioClose   io.Closer
		z         *zap.Logger
		TraceID   string
		ctxFields []string
	}
)

// 确保zapLogger实现了 ILogger 接口, 可以在编译时发现错误.
var _ ILogger = &ZapLogger{}

// NewLogger 初始化日志信息.
func NewLogger(opts ...Option) *ZapLogger {
	opt := newDefaultOptions()
	for _, f := range opts {
		f(opt)
	}

	// 设置默认的配置信息
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.MessageKey = "message"
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(opt.timeLayout))
	}
	// 指定 time Duration 序列化函数, 将time.Duration 序列化经过的毫秒数的浮点数
	// 毫秒数比默认的秒数更精确
	encoderConfig.EncodeDuration = func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendFloat64(float64(d) / float64(time.Millisecond))
	}

	jsonEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	cores := make([]zapcore.Core, 0, 3)
	level := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= opt.level
	})

	// 将日志输出到标准输出中
	if !opt.disableConsole {
		stdout := zapcore.Lock(os.Stdout)

		cores = append(cores,
			zapcore.NewCore(jsonEncoder, zapcore.NewMultiWriteSyncer(stdout), level),
		)
	}

	// 将日志输出到文件
	if opt.file != nil {
		cores = append(cores,
			zapcore.NewCore(jsonEncoder, zapcore.AddSync(opt.file), level),
		)
	}

	// 创建 *zap.Logger 对象
	zapOpts := make([]zap.Option, 0, 3)
	zapOpts = append(zapOpts, zap.AddCallerSkip(1))

	if !opt.disableCaller {
		zapOpts = append(zapOpts, zap.AddCaller())
	}

	if !opt.disableStacktrace {
		zapOpts = append(zapOpts, zap.AddStacktrace(zap.PanicLevel))
	}

	z := &ZapLogger{
		z:         zap.New(zapcore.NewTee(cores...), zapOpts...),
		ctxFields: opt.ctxFields,
		TraceID:   opt.traceID,
		ioClose:   opt.ioClose,
	}

	// 将标准库的 log.Logger 的info级别的输出重定向到 zap.Logger
	zap.RedirectStdLog(z.z)

	return z
}

// 自己克隆自己，防止污染.
func (z *ZapLogger) clone() *ZapLogger {
	lc := *z

	return &lc
}

// 输出日志信息.
func (z *ZapLogger) print(ctx context.Context, level zapcore.Level, msg string, vals ...any) {
	lc := z.clone()

	if val, ok := ctx.Value(lc.TraceID).(string); ok {
		msg = fmt.Sprintf("%-14s%s", val, msg)
	}

	// 设置日志打印字段
	fields := make([]zap.Field, 0, len(lc.ctxFields))
	for _, field := range lc.ctxFields {
		if val := ctx.Value(field); val != nil {
			fields = append(fields, zap.Any(field, val))
		}
	}
	lc.z = lc.z.With(fields...)

	sugar := lc.z.Sugar()
	switch level {
	case zapcore.DebugLevel:
		sugar.Debugf(msg, vals...)
	case zapcore.InfoLevel:
		sugar.Infof(msg, vals...)
	case zapcore.WarnLevel:
		sugar.Warnf(msg, vals...)
	case zapcore.ErrorLevel:
		sugar.Errorf(msg, vals...)
	case zapcore.PanicLevel:
		sugar.Panicf(msg, vals...)
	case zapcore.FatalLevel:
		sugar.Fatalf(msg, vals...)
	}
}

// LogMode log mode.
func (z *ZapLogger) LogMode(level logger.LogLevel) logger.Interface {
	return z.clone()
}

func (z *ZapLogger) Debug(ctx context.Context, msg string, keysAndValues ...any) {
	z.print(ctx, zapcore.DebugLevel, msg, keysAndValues...)
}

func (z *ZapLogger) Printf(ctx context.Context, msg string, keysAndValues ...any) {
	z.print(ctx, zapcore.InfoLevel, msg, keysAndValues...)
}

func (z *ZapLogger) Info(ctx context.Context, msg string, keysAndValues ...any) {
	z.print(ctx, zapcore.InfoLevel, msg, keysAndValues...)
}

func (z *ZapLogger) Warn(ctx context.Context, msg string, keysAndValues ...any) {
	z.print(ctx, zapcore.WarnLevel, msg, keysAndValues...)
}

func (z *ZapLogger) Error(ctx context.Context, msg string, keysAndValues ...any) {
	z.print(ctx, zapcore.ErrorLevel, msg, keysAndValues...)
}

func (z *ZapLogger) Panic(ctx context.Context, msg string, keysAndValues ...any) {
	z.print(ctx, zapcore.PanicLevel, msg, keysAndValues...)
}

func (z *ZapLogger) Fatal(ctx context.Context, msg string, keysAndValues ...any) {
	z.print(ctx, zapcore.FatalLevel, msg, keysAndValues...)
}

func (z *ZapLogger) Sync() {
	if z.ioClose != nil {
		_ = z.ioClose.Close()
	}
	_ = z.z.Sync()
}

// 复制gorm中log日志打印逻辑
//

func (z *ZapLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	z.Info(
		ctx,
		"耗时:%.2fs(ns: %.2fns) 行数:%d \n[sql] %s\n",
		elapsed.Seconds(),
		float64(elapsed.Nanoseconds())/1e6,
		rows,
		sql,
	)
}
