package dhtfx

import (
	"github.com/melkor217/bitmagnet/internal/config/configfx"
	"github.com/melkor217/bitmagnet/internal/protocol"
	"github.com/melkor217/bitmagnet/internal/protocol/dht/client"
	"github.com/melkor217/bitmagnet/internal/protocol/dht/ktable"
	"github.com/melkor217/bitmagnet/internal/protocol/dht/responder"
	"github.com/melkor217/bitmagnet/internal/protocol/dht/server"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"dht",
		configfx.NewConfigModule[server.Config]("dht_server", server.NewDefaultConfig()),
		fx.Provide(
			fx.Annotated{
				Name: "dht_node_id",
				Target: func() protocol.ID {
					return protocol.RandomNodeIDWithClientSuffix()
				},
			},
			client.New,
			ktable.New,
			responder.New,
			server.New,
		),
	)
}
