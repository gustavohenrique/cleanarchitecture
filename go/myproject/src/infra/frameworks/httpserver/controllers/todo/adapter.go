package todo

import (
	"{{ .ProjectName }}/src/domain/entities"
	"{{ .ProjectName }}/src/shared/datetime"
)

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (p *Adapter) ToJsonResponse(in entities.Todo) TodoJsonResponse {
	out := TodoJsonResponse{}
	out.ID = in.ID
	out.Title = in.Title
	out.IsDone = in.IsDone
	out.CreatedAt = datetime.ToString(in.CreatedAt)
	return out
}

func (p *Adapter) ToJsonResponses(in []entities.Todo) []TodoJsonResponse {
	var out []TodoJsonResponse
	for _, item := range in {
		out = append(out, p.ToJsonResponse(item))
	}
	return out
}
