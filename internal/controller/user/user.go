package user

import "gin-study/internal/service/user"

type Controller struct {
	UserService user.Server
}

// NewController 获取用户控制器.
func NewController(userService user.Server) *Controller {
	return &Controller{
		UserService: userService,
	}
}
