package options

import (
	"time"

	"gorm.io/gorm"

	"gin-study/pkg/db"
	"gin-study/pkg/logger"
)

type MariadbOption struct {
	Logger                *logger.ZapLogger `json:"-"                        mapstructure:"-"`
	Address               string            `json:"address"                  mapstructure:"address"`
	Username              string            `json:"username"                 mapstructure:"username"`
	Password              string            `json:"password"                 mapstructure:"password"`
	Database              string            `json:"database"                 mapstructure:"database"`
	MaxOpenConnections    int               `json:"max_open_connections"     mapstructure:"max_open_connections"`
	MaxIdleConnections    int               `json:"max_idle_connections"     mapstructure:"max_idle_connections"`
	MaxConnectionLifeTime time.Duration     `json:"max_connection_life_time" mapstructure:"max_connection_life_time"`
}

func (o *MariadbOption) Option() *db.Option {
	return &db.Option{
		Address:               o.Address,
		Username:              o.Username,
		Password:              o.Password,
		Database:              o.Database,
		MaxIdleConnections:    o.MaxIdleConnections,
		MaxOpenConnections:    o.MaxOpenConnections,
		MaxConnectionLifeTime: o.MaxConnectionLifeTime,
		Logger:                o.Logger,
	}
}

func (o *MariadbOption) NewClient(logger *logger.ZapLogger) (*gorm.DB, error) {
	o.Logger = logger

	return db.New(o.Option())
}
