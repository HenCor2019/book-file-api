package v1

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/HenCor2019/book-file-api/api/config"
	health_rts "github.com/HenCor2019/book-file-api/api/v1/health"
)

const (
	PREFIX  = "/api"
	VERSION = "v1"
)

type API struct {
	HealtRts health_rts.HealthRtr
	Logger   *slog.Logger
}

func New(
	healthRt health_rts.HealthRtr,
	logger *slog.Logger,
) *API {
	return &API{
		HealtRts: healthRt,
		Logger:   logger,
	}
}

func (api *API) Start(v1 *config.RouteBundle) error {
	PORT := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", PORT)

	api.Logger.Info(fmt.Sprintf("App is ready and listening on port %s 🚀", PORT))
	return http.ListenAndServe(addr, v1)
}
