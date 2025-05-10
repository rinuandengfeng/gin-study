package options

import (
	"github.com/redis/go-redis/v9"

	"gin-study/pkg/logger"
	"gin-study/pkg/mredis"
)

type RedisOption struct {
	Logger   *logger.ZapLogger `json:"-"        mapstructure:"-"`
	Address  string            `json:"address"  mapstructure:"address"`
	Password string            `json:"password" mapstructure:"password"`
	DB       int               `json:"db"       mapstructure:"db"`
}

func (o *RedisOption) Option() *mredis.Option {
	return &mredis.Option{
		Addr:     o.Address,
		Password: o.Password,
		DB:       o.DB,
		Logger:   o.Logger,
	}
}

func (o *RedisOption) NewClient(logger *logger.ZapLogger) (*redis.Client, error) {
	o.Logger = logger

	return mredis.New(o.Option())
}
