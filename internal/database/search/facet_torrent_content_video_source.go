package search

import (
	"github.com/melkor217/bitmagnet/internal/database/dao"
	"github.com/melkor217/bitmagnet/internal/database/query"
	"github.com/melkor217/bitmagnet/internal/model"
	"gorm.io/gen/field"
)

const VideoSourceFacetKey = "video_source"

func VideoSourceFacet(options ...query.FacetOption) query.Facet {
	return videoSourceFacet{torrentContentAttributeFacet[model.VideoSource]{
		FacetConfig: query.NewFacetConfig(
			append([]query.FacetOption{
				query.FacetHasKey(VideoSourceFacetKey),
				query.FacetHasLabel("Video Source"),
				query.FacetUsesOrLogic(),
				query.FacetTriggersCte(),
			}, options...)...,
		),
		field: func(q *dao.Query) field.Field {
			return q.TorrentContent.VideoSource
		},
		parse: model.ParseVideoSource,
	}}
}

type videoSourceFacet struct {
	torrentContentAttributeFacet[model.VideoSource]
}

func (videoSourceFacet) Values(query.FacetContext) (map[string]string, error) {
	vsrcs := model.VideoSourceValues()
	values := make(map[string]string, len(vsrcs))

	for _, vr := range vsrcs {
		values[vr.String()] = vr.Label()
	}

	return values, nil
}
