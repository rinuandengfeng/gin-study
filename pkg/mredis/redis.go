package mredis

import (
	"context"

	"github.com/redis/go-redis/v9"

	"gin-study/pkg/logger"
)

type Option struct {
	Logger   *logger.ZapLogger
	Addr     string
	Password string
	DB       int
}

func New(opt *Option) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     opt.Addr,
		Password: opt.Password,
		DB:       opt.DB,
	})

	// 添加日志
	if opt.Logger != nil {
		rdb.AddHook(NewLogHook(opt.Logger))
	}

	redis.SetLogger(opt.Logger)

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
