package queuefx

import (
	"github.com/melkor217/bitmagnet/internal/queue/manager"
	"github.com/melkor217/bitmagnet/internal/queue/prometheus"
	"github.com/melkor217/bitmagnet/internal/queue/server"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"queue",
		fx.Provide(
			server.New,
			manager.New,
			prometheus.New,
		),
	)
}
