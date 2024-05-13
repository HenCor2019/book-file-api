package services

type HealthCheckService interface {
	CheckHealth() string
	CheckHelloWorld() string
}

type Service struct {
}

func New() HealthCheckService {
	return &Service{}
}

func (s *Service) CheckHealth() string {
	return "ok"
}

func (s *Service) CheckHelloWorld() string {
	return "Hello World"
}
