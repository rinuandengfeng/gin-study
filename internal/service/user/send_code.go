package user

import "context"

const (
	SendCodeTypeLogin SendCodeType = "login"
)

type (
	SendCodeType string

	// SendCodeRequest 发送验证码请求信息.
	SendCodeRequest struct {
		Mobile string `binding:"required" form:"mobile" json:"mobile" label:"手机号"` // 手机号
		// @param type body string true "发送验证码类型" Enums(login, register) default(login)
		Type SendCodeType `binding:"required" form:"type"   json:"type"   label:"验证码类型"`
	}

	// SendCodeResponse 发送验证码返回信息.
	SendCodeResponse struct {
		Code string `json:"code"` // 验证码
	}
)

func (s *Service) SendCode(context.Context, *SendCodeRequest) (*SendCodeResponse, error) {
	// TODO: 待实现

	return &SendCodeResponse{}, nil
}
