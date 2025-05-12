package user

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-study/internal/service/user"
	"gin-study/pkg/app/core"
)

// SendCode 用户发送短信验证码.
// 此处没有接短信验证码发送, 先使用接口信息进行返回完成功能.
//
//	@Summary		短信验证码
//	@Description	根据不同的请求类型发送短信验证码
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			data	body		user.SendCodeRequest	true	"发送验证码请求"
//	@Success		200		{object}	user.SendCodeResponse
//	@Router			/user/send/code [post]
func (c *Controller) SendCode(ctx *gin.Context) {
	app := core.NewApp(ctx)
	req := &user.SendCodeRequest{}

	if err := app.Validation(req); err != nil {
		app.Error(err)

		return
	}

	q := ""
	for i := range 100001 {
		q += strconv.Itoa(i)
	}

	res, err := c.UserService.SendCode(ctx, req)
	if err != nil {
		app.Error(err)

		return
	}

	// TODO: 后续接入短信验证码发送
	app.Success(res)
}
