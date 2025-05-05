package databasefx

import (
	"github.com/melkor217/bitmagnet/internal/config/configfx"
	"github.com/melkor217/bitmagnet/internal/database"
	"github.com/melkor217/bitmagnet/internal/database/cache"
	"github.com/melkor217/bitmagnet/internal/database/dao"
	"github.com/melkor217/bitmagnet/internal/database/healthcheck"
	"github.com/melkor217/bitmagnet/internal/database/migrations"
	"github.com/melkor217/bitmagnet/internal/database/postgres"
	"github.com/melkor217/bitmagnet/internal/database/search"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"database",
		configfx.NewConfigModule[postgres.Config]("postgres", postgres.NewDefaultConfig()),
		configfx.NewConfigModule[cache.Config]("gorm_cache", cache.NewDefaultConfig()),
		fx.Provide(
			cache.NewInMemoryCacher,
			cache.NewPlugin,
			dao.New,
			database.New,
			healthcheck.New,
			migrations.New,
			postgres.New,
			search.New,
		),
		fx.Decorate(
			cache.NewDecorator,
		),
	)
}
