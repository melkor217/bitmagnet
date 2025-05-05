package tmdbhealthcheck

import (
	"context"
	"time"

	"github.com/melkor217/bitmagnet/internal/health"
	"github.com/melkor217/bitmagnet/internal/lazy"
	"github.com/melkor217/bitmagnet/internal/tmdb"
)

func NewCheck(
	enabled bool,
	client lazy.Lazy[tmdb.Client],
) health.Check {
	return health.Check{
		Name:    "tmdb",
		Timeout: time.Second * 30,
		IsActive: func() bool {
			return enabled
		},
		Check: func(ctx context.Context) error {
			c, err := client.Get()
			if err != nil {
				return err
			}
			return c.ValidateAPIKey(ctx)
		},
	}
}
