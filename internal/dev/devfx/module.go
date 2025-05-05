package devfx

import (
	"github.com/melkor217/bitmagnet/internal/app/cli"
	"github.com/melkor217/bitmagnet/internal/app/cli/args"
	"github.com/melkor217/bitmagnet/internal/config/configfx"
	"github.com/melkor217/bitmagnet/internal/database"
	"github.com/melkor217/bitmagnet/internal/database/migrations"
	"github.com/melkor217/bitmagnet/internal/database/postgres"
	"github.com/melkor217/bitmagnet/internal/dev/app/cmd/gormcmd"
	"github.com/melkor217/bitmagnet/internal/dev/app/cmd/migratecmd"
	"github.com/melkor217/bitmagnet/internal/logging/loggingfx"
	"github.com/melkor217/bitmagnet/internal/validation/validationfx"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"dev",
		configfx.NewConfigModule[postgres.Config]("postgres", postgres.NewDefaultConfig()),
		configfx.New(),
		loggingfx.New(),
		validationfx.New(),
		fx.Provide(args.New),
		fx.Provide(cli.New),
		fx.Provide(database.New),
		fx.Provide(migrations.New),
		fx.Provide(postgres.New),
		fx.Provide(gormcmd.New),
		fx.Provide(migratecmd.New),
	)
}
