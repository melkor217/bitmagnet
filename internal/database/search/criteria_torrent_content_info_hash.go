package search

import (
	"github.com/melkor217/bitmagnet/internal/database/query"
	"github.com/melkor217/bitmagnet/internal/model"
	"github.com/melkor217/bitmagnet/internal/protocol"
)

func TorrentContentInfoHashCriteria(infoHashes ...protocol.ID) query.Criteria {
	return infoHashCriteria(model.TableNameTorrentContent, infoHashes...)
}
