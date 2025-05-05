package app

import (
	"github.com/melkor217/bitmagnet/internal/app/appfx"
	"github.com/melkor217/bitmagnet/internal/app/cli/hooks"
	"github.com/melkor217/bitmagnet/internal/logging/loggingfx"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func New() *fx.App {
	return fx.New(
		appfx.New(),
		loggingfx.WithLogger(),
		fx.Invoke(func(
			logger *zap.SugaredLogger,
			_ *cli.App,
			_ hooks.AttachedHooks,
		) {
			logger.Debug("app invoked")
		}),
	)
}
