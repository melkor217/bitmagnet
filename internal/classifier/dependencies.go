package classifier

import (
	"github.com/melkor217/bitmagnet/internal/tmdb"
)

type dependencies struct {
	search     LocalSearch
	tmdbClient tmdb.Client
}
