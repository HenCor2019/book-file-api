package controllers

import (
	"net/http"

	"github.com/HenCor2019/book-file-api/internal/health/services"
)

type HealthCheckController interface {
	HealthCheckHandler(w http.ResponseWriter, r *http.Request)
	HelloWorldHandler(w http.ResponseWriter, r *http.Request)
}

type Controller struct {
	healthCheckService services.HealthCheckService
}

func New(healthCheckService services.HealthCheckService) HealthCheckController {
	return &Controller{healthCheckService: healthCheckService}
}

func (c *Controller) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	healthStatus := c.healthCheckService.CheckHealth()

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(healthStatus)); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (c *Controller) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	helloWorldMessage := c.healthCheckService.HelloWorld()
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(helloWorldMessage)); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
