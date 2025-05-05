package torznabfx

import (
	"github.com/melkor217/bitmagnet/internal/config/configfx"
	"github.com/melkor217/bitmagnet/internal/database/search"
	"github.com/melkor217/bitmagnet/internal/lazy"
	"github.com/melkor217/bitmagnet/internal/torznab"
	"github.com/melkor217/bitmagnet/internal/torznab/adapter"
	"github.com/melkor217/bitmagnet/internal/torznab/httpserver"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"torznab",
		configfx.NewConfigModule[torznab.Config]("torznab", torznab.NewDefaultConfig()),
		fx.Provide(
			func(lazySearch lazy.Lazy[search.Search]) lazy.Lazy[torznab.Client] {
				return lazy.New[torznab.Client](func() (torznab.Client, error) {
					s, err := lazySearch.Get()
					if err != nil {
						return nil, err
					}

					return adapter.New(s), nil
				})
			},
			fx.Annotate(
				httpserver.New,
				fx.ResultTags(`group:"http_server_options"`),
			),
		),
		fx.Decorate(
			func(cfg torznab.Config) torznab.Config {
				return cfg.MergeDefaults()
			}),
	)
}
