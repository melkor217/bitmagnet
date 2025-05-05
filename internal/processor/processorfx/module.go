package processorfx

import (
	"github.com/melkor217/bitmagnet/internal/config/configfx"
	"github.com/melkor217/bitmagnet/internal/processor"
	batchqueue "github.com/melkor217/bitmagnet/internal/processor/batch/queue"
	processorqueue "github.com/melkor217/bitmagnet/internal/processor/queue"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"processor",
		configfx.NewConfigModule[processor.Config]("processor", processor.NewDefaultConfig()),
		fx.Provide(
			processor.New,
			processorqueue.New,
			batchqueue.New,
		),
	)
}
