package core

import "net/http"

const (
	// 通用错误.
	UnKnownErrCode      = 1
	SystemErrCode       = 2
	NotFoundErrCode     = 100006
	BadRequestErrCode   = 100007
	UnauthorizedErrCode = 100008

	// 用户业务错误.
	UserAccountPasswordErrCode   = 101001
	UserLoginSendRepeatErrCode   = 101002
	UserLoginCodeValidateErrCode = 101003
)

var (
	// 错误code码对应的消息提示.
	ErrorMsg = map[int]string{
		UnKnownErrCode:               "未知错误",
		SystemErrCode:                "系统错误",
		NotFoundErrCode:              "页面不存在",
		BadRequestErrCode:            "请求参数错误",
		UnauthorizedErrCode:          "鉴权失败，请重新登录",
		UserAccountPasswordErrCode:   "账号密码错误",
		UserLoginSendRepeatErrCode:   "验证码重复发送",
		UserLoginCodeValidateErrCode: "验证码错误",
	}

	// 错误code码对应的解决方案提示.
	ErrorRef = make(map[int]string, 0)

	ErrUnknown = New(
		http.StatusInternalServerError,
		UnKnownErrCode,
		ErrorMsg[UnKnownErrCode],
		ErrorRef[UnKnownErrCode],
	)
	ErrSystem = New(
		http.StatusInternalServerError,
		SystemErrCode,
		ErrorMsg[SystemErrCode],
		ErrorRef[SystemErrCode],
	)
	ErrNotFound = New(
		http.StatusNotFound,
		NotFoundErrCode,
		ErrorMsg[NotFoundErrCode],
		ErrorRef[NotFoundErrCode],
	)
	ErrBadRequest = New(
		http.StatusOK,
		BadRequestErrCode,
		ErrorMsg[BadRequestErrCode],
		ErrorRef[BadRequestErrCode],
	)
	ErrUnauthorized = New(
		http.StatusOK,
		UnauthorizedErrCode,
		ErrorMsg[UnauthorizedErrCode],
		ErrorRef[UnauthorizedErrCode],
	)
	// 用户业务相关错误.
	ErrUserAccountPassword = New(
		http.StatusOK,
		UserAccountPasswordErrCode,
		ErrorMsg[UserAccountPasswordErrCode],
		ErrorRef[UserAccountPasswordErrCode],
	)
	ErrUserLoginSendRepeat = New(
		200,
		UserLoginSendRepeatErrCode,
		ErrorMsg[UserLoginSendRepeatErrCode],
		ErrorRef[UserLoginSendRepeatErrCode],
	)
	ErrUserLoginCodeValidate = New(
		http.StatusOK,
		UserLoginCodeValidateErrCode,
		ErrorMsg[UserLoginCodeValidateErrCode],
		ErrorRef[UserLoginCodeValidateErrCode],
	)
)
