package importerfx

import (
	"github.com/melkor217/bitmagnet/internal/importer"
	"github.com/melkor217/bitmagnet/internal/importer/httpserver"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"importer",
		fx.Provide(
			httpserver.New,
			importer.New,
		),
	)
}
