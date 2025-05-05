package healthfx

import (
	"github.com/melkor217/bitmagnet/internal/health"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"health",
		fx.Provide(
			health.New,
		),
	)
}
