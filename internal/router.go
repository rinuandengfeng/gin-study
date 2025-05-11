package internal

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "gin-study/api"
	"gin-study/internal/controller/user"
	"gin-study/internal/middleware"
	userServ "gin-study/internal/service/user"
)

// 初始化 gin 路由.
func initRouter(ctx context.Context, apiServ *apiServer) {
	r := gin.New()
	gin.SetMode(apiServ.option.Server.RunMode)

	registerRouter(r, apiServ)

	server := &http.Server{
		Addr:           apiServ.option.Server.Address,
		Handler:        r,
		ReadTimeout:    apiServ.option.Server.ReadTimeout,
		WriteTimeout:   apiServ.option.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	//nolint: contextcheck
	apiServ.ShutDown.Register(func() {
		nCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		apiServ.Logger.Info(nCtx, "正在关闭服务")
		if err := server.Shutdown(nCtx); err != nil {
			apiServ.Logger.Error(nCtx, "服务被迫关闭：%s", err.Error())
		}

		apiServ.Logger.Info(nCtx, "http服务已关闭")
	})
}

func registerRouter(r *gin.Engine, apiServ *apiServer) {
	if apiServ.option.Server.RunMode == gin.DebugMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Use(middleware.Recovery(apiServ.Logger), middleware.Cors(), middleware.RequestLog(apiServ.Logger))

	apiv1 := r.Group("/v1")
	{
		userv1 := apiv1.Group("/user")
		{
			userCon := user.NewController(userServ.NewService())
			userv1.POST("/send/code", userCon.SendCode)
		}
	}
}
