package user

// Service 用户服务实现.
type Service struct{}

func NewService() Server {
	return &Service{}
}
