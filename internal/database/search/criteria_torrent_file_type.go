package search

import (
	"github.com/melkor217/bitmagnet/internal/database/query"
	"github.com/melkor217/bitmagnet/internal/model"
)

func TorrentFileTypeCriteria(fileTypes ...model.FileType) query.Criteria {
	var extensions []string
	for _, fileType := range fileTypes {
		extensions = append(extensions, fileType.Extensions()...)
	}

	return TorrentFileExtensionCriteria(extensions...)
}
