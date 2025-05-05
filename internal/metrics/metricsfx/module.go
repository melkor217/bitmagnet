package metricsfx

import (
	"github.com/melkor217/bitmagnet/internal/metrics/queuemetrics"
	"github.com/melkor217/bitmagnet/internal/metrics/torrentmetrics"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"queue",
		fx.Provide(
			queuemetrics.New,
			torrentmetrics.New,
		),
	)
}
