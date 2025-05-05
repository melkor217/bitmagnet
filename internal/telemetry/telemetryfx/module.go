package telemetryfx

import (
	"github.com/melkor217/bitmagnet/internal/telemetry/httpserver"
	"github.com/melkor217/bitmagnet/internal/telemetry/prometheus"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"telemetry",
		fx.Provide(
			httpserver.New,
			prometheus.New,
		),
	)
}
