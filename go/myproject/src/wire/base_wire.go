package wire

import "{{ .ProjectName }}/src/domain/models"

type TodoHttpRequest struct {
	ID string `json:"id"`
}

func (wireIn *TodoHttpRequest) ToModel() models.TodoModel {
	return models.TodoModel{}
}

func (wireIn *TodoHttpRequest) Validate() error {
	return nil
}

type TodoHttpResponse struct {
	ID string `json:"id"`
}

func (wireOut TodoHttpResponse) Of(model models.TodoModel) TodoHttpResponse {
	wireOut.ID = model.ID
	return wireOut
}

type TodoEntity struct {
	ID string
}

func (entity TodoEntity) ToModel() models.TodoModel {
	return models.TodoModel{}
}

func (entity TodoEntity) Of(m models.TodoModel) TodoEntity {
	return entity
}
