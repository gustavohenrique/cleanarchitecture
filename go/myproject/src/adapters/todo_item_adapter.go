package adapters

import (
	"myproject/src/entities"
	pb "myproject/src/proto"
	"myproject/src/valueobjects"
	"myproject/src/shared/datetime"
)

type TodoItemAdapter struct{}

func NewTodoItemAdapter() TodoItemAdapter {
	return TodoItemAdapter{}
}

func (a TodoItemAdapter) FromEntityToProto(in entities.TodoItemEntity) *pb.TodoItem {
	out := &pb.TodoItem{}
	out.Id = in.ID
	out.CreatedAt = datetime.ToString(in.CreatedAt)
	out.Title = in.Title
	out.IsDone = in.IsDone
	return out
}

func (a TodoItemAdapter) FromEntityListToProtoList(in []entities.TodoItemEntity) []*pb.TodoItem {
	var out []*pb.TodoItem
	for _, i := range in {
		out = append(out, a.FromEntityToProto(i))
	}
	return out
}

func (a TodoItemAdapter) FromProtoToEntity(in *pb.TodoItem) entities.TodoItemEntity {
	out := entities.TodoItemEntity{}
	if in == nil {
		return out
	}
	out.ID = in.Id
	out.Title = in.Title
	out.IsDone = in.IsDone
	return out
}

func (a TodoItemAdapter) FromTableToEntity(in valueobjects.TodoItemTable) entities.TodoItemEntity {
	out := entities.TodoItemEntity{}
	out.ID = in.ID
	out.Title = in.Title
	out.CreatedAt = datetime.Parse(in.CreatedAt).ToISO()
	out.IsDone = in.IsDone
	return out
}

func (a TodoItemAdapter) FromTableListToEntityList(in []valueobjects.TodoItemTable) []entities.TodoItemEntity {
	var out []entities.TodoItemEntity
	for _, item := range in {
		out = append(out, a.FromTableToEntity(item))
	}
	return out
}

func (a TodoItemAdapter) FromRequestToEntity(in valueobjects.TodoItemRequest) entities.TodoItemEntity {
	out := entities.TodoItemEntity{}
	out.ID = in.ID
	out.Title = in.Title
	out.IsDone = in.IsDone
	return out
}

func (a TodoItemAdapter) FromEntityToResponse(in entities.TodoItemEntity) valueobjects.TodoItemResponse {
	out := valueobjects.TodoItemResponse{}
	out.ID = in.ID
	out.Title = in.Title
	out.IsDone = in.IsDone
	out.CreatedAt = datetime.ToString(in.CreatedAt)
	return out
}

func (a TodoItemAdapter) FromEntityListToResponseList(in []entities.TodoItemEntity) []valueobjects.TodoItemResponse {
	var out []valueobjects.TodoItemResponse
	for _, item := range in {
		out = append(out, a.FromEntityToResponse(item))
	}
	return out
}
