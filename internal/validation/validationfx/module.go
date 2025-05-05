package validationfx

import (
	"github.com/melkor217/bitmagnet/internal/validation"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"validation",
		fx.Provide(validation.New),
	)
}
