package user

import "context"

// Server 用户服务对外暴露接口.
type Server interface {
	SendCode(context.Context, *SendCodeRequest) (*SendCodeResponse, error)
}
