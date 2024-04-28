package services

import "log"

type HealthCheckService interface {
	CheckHealth() string
}

type Service struct {
}

func New() HealthCheckService {
	return &Service{}
}

func (s *Service) CheckHealth() string {
	log.Println("GET healthcheck")
	return "ok"
}
