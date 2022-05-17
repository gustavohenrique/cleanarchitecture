package adapters

import (
	"context"

	"myproject/src/interfaces"
	pb "myproject/src/proto"
	"myproject/src/valueobjects/sqlite"
)

type SearchAdapter struct{}

func NewSearchAdapter() SearchAdapter {
	return SearchAdapter{}
}

func (a SearchAdapter) FromProtoToVO(ctx context.Context, in *pb.SearchRequest) interfaces.ISearchRequest {
	v := ctx.Value("db")
	if v == nil {
		return sqlite.SearchRequest{}
	}
	out := sqlite.SearchRequest{}
	if in == nil {
		return out
	}

	out.Pagination = sqlite.Pagination{
		Page:    in.Page,
		PerPage: in.PerPage,
	}

	return out
}
