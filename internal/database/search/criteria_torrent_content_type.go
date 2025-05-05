package search

import (
	"github.com/melkor217/bitmagnet/internal/database/query"
	"github.com/melkor217/bitmagnet/internal/maps"
	"github.com/melkor217/bitmagnet/internal/model"
	"gorm.io/gen/field"
)

func TorrentContentTypeCriteria(types ...model.ContentType) query.Criteria {
	strTypes := make([]string, 0, len(types))
	for _, contentType := range types {
		strTypes = append(strTypes, contentType.String())
	}

	return query.DaoCriteria{
		Conditions: func(ctx query.DBContext) ([]field.Expr, error) {
			q := ctx.Query()
			return []field.Expr{
				q.TorrentContent.ContentType.In(strTypes...),
			}, nil
		},
		Joins: maps.NewInsertMap(
			maps.MapEntry[string, struct{}]{Key: model.TableNameTorrentContent},
		),
	}
}
