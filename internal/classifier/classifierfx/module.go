package classifierfx

import (
	"github.com/melkor217/bitmagnet/internal/classifier"
	"github.com/melkor217/bitmagnet/internal/config/configfx"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"workflow",
		configfx.NewConfigModule[classifier.Config]("classifier", classifier.NewDefaultConfig()),
		fx.Provide(
			classifier.New,
		),
	)
}
