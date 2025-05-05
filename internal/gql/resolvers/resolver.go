package resolvers

import (
	"github.com/melkor217/bitmagnet/internal/blocking"
	"github.com/melkor217/bitmagnet/internal/database/dao"
	"github.com/melkor217/bitmagnet/internal/database/search"
	"github.com/melkor217/bitmagnet/internal/health"
	"github.com/melkor217/bitmagnet/internal/metrics/queuemetrics"
	"github.com/melkor217/bitmagnet/internal/metrics/torrentmetrics"
	"github.com/melkor217/bitmagnet/internal/processor"
	"github.com/melkor217/bitmagnet/internal/queue/manager"
	"github.com/melkor217/bitmagnet/internal/worker"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Dao                  *dao.Query
	Search               search.Search
	Workers              worker.Registry
	Checker              health.Checker
	QueueMetricsClient   queuemetrics.Client
	QueueManager         manager.Manager
	TorrentMetricsClient torrentmetrics.Client
	Processor            processor.Processor
	BlockingManager      blocking.Manager
}
