package workerfx

import (
	"github.com/melkor217/bitmagnet/internal/worker"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"worker",
		fx.Provide(worker.NewRegistry),
	)
}
