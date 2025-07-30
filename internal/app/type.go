package app

type IService interface {
	PingPong
}

type Service struct{}

func New() IService {
	return &Service{}
}
