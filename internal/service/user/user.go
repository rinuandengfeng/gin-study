package user

import "context"

// -self_package: 生成的代码包的完成包导入路径，防止生成的代码中出现循环导入
// -destination: 生成的代码文件
// -package: 生成的mock类的包
//go:generate mockgen -self_package=gin-study/internal/service/user -destination mock_service.go -package user gin-study/internal/service/user Server

// Server 用户服务对外暴露接口.
type Server interface {
	SendCode(context.Context, *SendCodeRequest) (*SendCodeResponse, error)
}
