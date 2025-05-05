package search

import (
	"github.com/melkor217/bitmagnet/internal/database/query"
	"github.com/melkor217/bitmagnet/internal/maps"
	"github.com/melkor217/bitmagnet/internal/model"
	"gorm.io/gen"
)

func TorrentFileExtensionCriteria(extensions ...string) query.Criteria {
	return query.GenCriteria(func(ctx query.DBContext) (query.Criteria, error) {
		q := ctx.Query()

		return query.OrCriteria{
			Criteria: []query.Criteria{
				query.RawCriteria{
					Query: q.Torrent.Where(
						q.Torrent.Extension.In(extensions...),
					),
					Joins: maps.NewInsertMap(
						maps.MapEntry[string, struct{}]{Key: model.TableNameTorrent},
					),
				},
				query.RawCriteria{
					Query: gen.Exists(
						q.TorrentFile.Where(
							q.TorrentFile.InfoHash.EqCol(q.Torrent.InfoHash),
							q.TorrentFile.Extension.In(extensions...),
						),
					),
					Joins: maps.NewInsertMap(
						maps.MapEntry[string, struct{}]{Key: model.TableNameTorrent},
					),
				},
			},
		}, nil
	})
}
