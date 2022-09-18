package controllers

import (
	"time"

	"{{ .ProjectName }}/src/domain/entities"
	pb "{{ .ProjectName }}/src/proto"
)

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) ToProto(in entities.Todo) *pb.TodoItem {
	out := &pb.TodoItem{}
	out.Id = in.ID
	out.Title = in.Title
	out.IsDone = in.IsDone
	if in.CreatedAt != nil {
		out.CreatedAt = in.CreatedAt.Format(time.RFC3339)
	}
	return out
}
func (a *Adapter) ToProtos(in []entities.Todo) []*pb.TodoItem {
	var out []*pb.TodoItem
	for _, i := range in {
		out = append(out, a.ToProto(i))
	}
	return out
}

func (a *Adapter) ToEntity(in *pb.TodoItem) entities.Todo {
	out := entities.Todo{}
	if in == nil {
		return out
	}
	out.ID = in.Id
	out.Title = in.Title
	out.IsDone = in.IsDone
	return out
}
