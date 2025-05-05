package appfx

import (
	"github.com/melkor217/bitmagnet/internal/app/cli"
	"github.com/melkor217/bitmagnet/internal/app/cli/args"
	"github.com/melkor217/bitmagnet/internal/app/cli/hooks"
	"github.com/melkor217/bitmagnet/internal/app/cmd/classifiercmd"
	"github.com/melkor217/bitmagnet/internal/app/cmd/configcmd"
	"github.com/melkor217/bitmagnet/internal/app/cmd/processcmd"
	"github.com/melkor217/bitmagnet/internal/app/cmd/reprocesscmd"
	"github.com/melkor217/bitmagnet/internal/app/cmd/workercmd"
	"github.com/melkor217/bitmagnet/internal/blocking/blockingfx"
	"github.com/melkor217/bitmagnet/internal/classifier/classifierfx"
	"github.com/melkor217/bitmagnet/internal/config/configfx"
	"github.com/melkor217/bitmagnet/internal/database/databasefx"
	"github.com/melkor217/bitmagnet/internal/database/migrations"
	"github.com/melkor217/bitmagnet/internal/dhtcrawler/dhtcrawlerfx"
	"github.com/melkor217/bitmagnet/internal/gql/gqlfx"
	"github.com/melkor217/bitmagnet/internal/health/healthfx"
	"github.com/melkor217/bitmagnet/internal/httpserver/httpserverfx"
	"github.com/melkor217/bitmagnet/internal/importer/importerfx"
	"github.com/melkor217/bitmagnet/internal/logging/loggingfx"
	"github.com/melkor217/bitmagnet/internal/metrics/metricsfx"
	"github.com/melkor217/bitmagnet/internal/processor/processorfx"
	"github.com/melkor217/bitmagnet/internal/protocol/dht/dhtfx"
	"github.com/melkor217/bitmagnet/internal/protocol/metainfo/metainfofx"
	"github.com/melkor217/bitmagnet/internal/queue/queuefx"
	"github.com/melkor217/bitmagnet/internal/telemetry/telemetryfx"
	"github.com/melkor217/bitmagnet/internal/tmdb/tmdbfx"
	"github.com/melkor217/bitmagnet/internal/torznab/torznabfx"
	"github.com/melkor217/bitmagnet/internal/validation/validationfx"
	"github.com/melkor217/bitmagnet/internal/version/versionfx"
	"github.com/melkor217/bitmagnet/internal/webui"
	"github.com/melkor217/bitmagnet/internal/worker/workerfx"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"app",
		blockingfx.New(),
		classifierfx.New(),
		configfx.New(),
		dhtcrawlerfx.New(),
		dhtfx.New(),
		databasefx.New(),
		gqlfx.New(),
		healthfx.New(),
		httpserverfx.New(),
		importerfx.New(),
		loggingfx.New(),
		metainfofx.New(),
		metricsfx.New(),
		processorfx.New(),
		queuefx.New(),
		telemetryfx.New(),
		tmdbfx.New(),
		torznabfx.New(),
		validationfx.New(),
		versionfx.New(),
		workerfx.New(),
		fx.Provide(
			args.New,
			cli.New,
			hooks.New,
			// cli commands:
			classifiercmd.New,
			configcmd.New,
			reprocesscmd.New,
			processcmd.New,
			workercmd.New,
		),
		fx.Provide(webui.New),
		fx.Decorate(migrations.NewDecorator),
	)
}
