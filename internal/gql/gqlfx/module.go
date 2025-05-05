package gqlfx

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/melkor217/bitmagnet/internal/blocking"
	"github.com/melkor217/bitmagnet/internal/database/dao"
	"github.com/melkor217/bitmagnet/internal/database/search"
	"github.com/melkor217/bitmagnet/internal/gql"
	"github.com/melkor217/bitmagnet/internal/gql/config"
	"github.com/melkor217/bitmagnet/internal/gql/httpserver"
	"github.com/melkor217/bitmagnet/internal/gql/resolvers"
	"github.com/melkor217/bitmagnet/internal/health"
	"github.com/melkor217/bitmagnet/internal/lazy"
	"github.com/melkor217/bitmagnet/internal/metrics/queuemetrics"
	"github.com/melkor217/bitmagnet/internal/metrics/torrentmetrics"
	"github.com/melkor217/bitmagnet/internal/processor"
	"github.com/melkor217/bitmagnet/internal/queue/manager"
	"github.com/melkor217/bitmagnet/internal/worker"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"graphql",
		fx.Provide(
			config.New,
			httpserver.New,
			func(
				lcfg lazy.Lazy[gql.Config],
			) lazy.Lazy[graphql.ExecutableSchema] {
				return lazy.New(func() (graphql.ExecutableSchema, error) {
					cfg, err := lcfg.Get()
					if err != nil {
						return nil, err
					}

					return gql.NewExecutableSchema(cfg), nil
				})
			},
		),
		fx.Provide(
			func(p Params) Result {
				return Result{
					Resolver: lazy.New(func() (*resolvers.Resolver, error) {
						ch, err := p.Checker.Get()
						if err != nil {
							return nil, err
						}
						s, err := p.Search.Get()
						if err != nil {
							return nil, err
						}
						d, err := p.Dao.Get()
						if err != nil {
							return nil, err
						}
						qmc, err := p.QueueMetricsClient.Get()
						if err != nil {
							return nil, err
						}
						qm, err := p.QueueManager.Get()
						if err != nil {
							return nil, err
						}
						tm, err := p.TorrentMetricsClient.Get()
						if err != nil {
							return nil, err
						}
						pr, err := p.Processor.Get()
						if err != nil {
							return nil, err
						}
						bm, err := p.BlockingManager.Get()
						if err != nil {
							return nil, err
						}
						return &resolvers.Resolver{
							Dao:                  d,
							Search:               s,
							Checker:              ch,
							QueueMetricsClient:   qmc,
							QueueManager:         qm,
							TorrentMetricsClient: tm,
							Processor:            pr,
							BlockingManager:      bm,
						}, nil
					}),
				}
			},
		),
		// inject resolver dependencies avoiding a circular dependency:
		fx.Invoke(func(
			resolver lazy.Lazy[*resolvers.Resolver],
			workers worker.Registry,
		) {
			resolver.Decorate(func(r *resolvers.Resolver) (*resolvers.Resolver, error) {
				r.Workers = workers
				return r, nil
			})
		}),
	)
}

type Params struct {
	fx.In
	Search               lazy.Lazy[search.Search]
	Dao                  lazy.Lazy[*dao.Query]
	Checker              lazy.Lazy[health.Checker]
	QueueMetricsClient   lazy.Lazy[queuemetrics.Client]
	QueueManager         lazy.Lazy[manager.Manager]
	TorrentMetricsClient lazy.Lazy[torrentmetrics.Client]
	Processor            lazy.Lazy[processor.Processor]
	BlockingManager      lazy.Lazy[blocking.Manager]
}

type Result struct {
	fx.Out
	Resolver lazy.Lazy[*resolvers.Resolver]
}
