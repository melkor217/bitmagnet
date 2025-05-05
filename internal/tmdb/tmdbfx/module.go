package tmdbfx

import (
	"github.com/melkor217/bitmagnet/internal/config/configfx"
	"github.com/melkor217/bitmagnet/internal/tmdb"
	"github.com/melkor217/bitmagnet/internal/tmdb/tmdbhealthcheck"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"tmdb",
		configfx.NewConfigModule[tmdb.Config]("tmdb", tmdb.NewDefaultConfig()),
		fx.Provide(
			tmdb.New,
			tmdbhealthcheck.New,
		),
	)
}
