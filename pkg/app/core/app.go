package core

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type PageReq struct {
//	Limit  int `form:"limit"`
//	Offset int `form:"offset"`
//}

type App struct {
	Ctx *gin.Context
}

func NewApp(ctx *gin.Context) *App {
	resp := &App{Ctx: ctx}

	return resp
}

func (a *App) Success(data any) {
	if data == nil {
		data = gin.H{}
	}

	a.Ctx.JSON(http.StatusOK, data)
}

// Error 返回错误消息.
func (a *App) Error(err error) {
	var appErr *Error
	if !errors.As(err, &appErr) {
		appErr = ErrUnknown
	}

	a.Ctx.AbortWithStatusJSON(appErr.HTTP(), gin.H{
		"code": appErr.Code(),
		"msg":  appErr.Msg(),
		"ref":  appErr.Ref(),
	})
}

func (a *App) Validation(obj any) error {
	if err := a.Ctx.ShouldBind(obj); err != nil {
		return ErrBadRequest.SetMsg(err.Error())
	}

	return nil
}
