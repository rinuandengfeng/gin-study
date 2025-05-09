package user

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"gin-study/internal/pkg/core"
	"gin-study/internal/pkg/options"
	"gin-study/pkg/logger"
	"gin-study/pkg/utils"
)

type (
	// 项目启动需要用到得配置选项.
	option struct {
		Server  *options.ServerOption  `mapstructure:"server"`
		Logger  *options.LoggerOption  `mapstructure:"logger"`
		Mariadb *options.MariadbOption `mapstructure:"mariadb"`
		Redis   *options.RedisOption   `mapstructure:"redis"`
	}

	apiServer struct {
		// 服务关闭时调用
		ShutDown utils.IShutDown

		Logger *logger.ZapLogger
		DB     *gorm.DB
		Redis  *redis.Client

		option *option
	}
)

// api服务启动初始化.
func newApiServer(conf string) (*apiServer, error) {
	apiServ := &apiServer{
		option: &option{},

		ShutDown: utils.NewShutDown(),
	}

	// 从配置文件中加载配置项
	vp := viper.New()
	vp.SetConfigFile(conf)
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := vp.Unmarshal(apiServ.option); err != nil {
		return nil, err
	}

	if err := apiServ.init(); err != nil {
		return nil, err
	}

	return apiServ, nil
}

// 初始化使用工具.
func (apiServ *apiServer) init() error {
	apiServ.option.Logger.TraceID = core.HeaderKeyTraceID
	apiServ.option.Logger.Fields = []string{
		core.HeaderKeyAuthorize,
	}
	apiServ.Logger = apiServ.option.Logger.NewClient()
	var err error

	if apiServ.DB, err = apiServ.option.Mariadb.NewClient(apiServ.Logger); err != nil {
		return err
	}

	if apiServ.Redis, err = apiServ.option.Redis.NewClient(apiServ.Logger); err != nil {
		return err
	}

	return nil
}

// Run 启动服务.
func (apiServ *apiServer) Run(ctx context.Context) error {
	initRouter(ctx, apiServ)
	apiServ.Close(ctx)

	return nil
}

// Close 服务关闭.
func (apiServ *apiServer) Close(ctx context.Context) {
	apiServ.ShutDown.Close()
}
