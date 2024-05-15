package routes

import (
	"github.com/HenCor2019/book-file-api/api/config"
	health_ctnlr "github.com/HenCor2019/book-file-api/internal/health/controllers"
)

type HealthRtr interface{}

type Rtr struct {
	health_ctnlr health_ctnlr.HealthCheckController
	router       *config.RouteBundle
}

func New(health_ctnlr health_ctnlr.HealthCheckController, router *config.RouteBundle) HealthRtr {
	healthGroup := router.Group()
	healthGroup.HandleFunc("GET /healthcheck", health_ctnlr.HealthCheckHandler)
	healthGroup.HandleFunc("GET /hello-world", health_ctnlr.HelloWorldHandler)
	return &Rtr{health_ctnlr, router}
}
