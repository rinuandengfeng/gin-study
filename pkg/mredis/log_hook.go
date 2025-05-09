package mredis

import (
	"context"
	"net"
	"time"

	"github.com/redis/go-redis/v9"

	"gin-study/pkg/logger"
)

type LogHook struct {
	logger logger.ILogger
}

// NewLogHook 自定义 redis 日志钩子， 需要实现 redis.Hook 接口.
func NewLogHook(logger logger.ILogger) *LogHook {
	return &LogHook{
		logger: logger,
	}
}

// DialHook 创建网络连接时调用的hook.
func (l *LogHook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		l.logger.Info(ctx, "调用DialHook")

		return next(ctx, network, addr)
	}
}

// ProcessHook 执行命令时调用的hook.
func (l *LogHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		begin := time.Now()
		_ = next(ctx, cmd)
		elapsed := time.Since(begin)
		l.logger.Info(
			ctx,
			"耗时：%.2fs(ns: %.2fns) \n[redis] %s\n",
			elapsed.Seconds(),
			float64(elapsed.Nanoseconds())/1e6,
			cmd.String(),
		)

		return nil
	}
}

// ProcessPipelineHook 执行管道命令时调用的hook.
func (l *LogHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		l.logger.Info(ctx, "调用ProcessPipelineHook方法")

		return nil
	}
}
