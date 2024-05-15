package services

import (
	"fmt"
	"os"
)

type HealthCheckService interface {
	CheckHealth() string
	HelloWorld() string
}

type Service struct {
}

func New() HealthCheckService {
	return &Service{}
}

func (s *Service) CheckHealth() string {
	return "ok"
}

func (s *Service) HelloWorld() string {
	env := os.Getenv("ENV")
	return fmt.Sprintf("Hello World from ENV: %s", env)
}
