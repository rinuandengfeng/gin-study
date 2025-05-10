package core

type Error struct {
	msg  string
	ref  string
	http int
	code int
}

func (e *Error) Error() string { return e.msg }
func (e *Error) HTTP() int     { return e.http }
func (e *Error) Code() int     { return e.code }
func (e *Error) Msg() string   { return e.msg }
func (e *Error) Ref() string   { return e.ref }

func New(http, code int, msg, ref string) *Error {
	return &Error{http: http, code: code, msg: msg, ref: ref}
}

func (e *Error) copy() *Error {
	ec := *e

	return &ec
}

// SetMsg 修改返回错误信息.
func (e *Error) SetMsg(msg string) *Error {
	ec := e.copy()
	ec.msg = msg

	return ec
}
