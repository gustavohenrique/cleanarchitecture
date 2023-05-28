package usecases

import (
	"context"

	"{{ .ProjectName }}/src/domain/models"
	"{{ .ProjectName }}/src/domain/ports"
)

{{ range .Models }}
type Create{{ .CamelCaseName }}UseCase struct {
	{{ .LowerCaseName }}Repository ports.{{ .CamelCaseName }}Repository
}

func NewCreate{{ .CamelCaseName }}UseCase(repos ports.Repositories) Create{{ .CamelCaseName }}UseCase {
	return Create{{ .CamelCaseName }}UseCase{
		{{ .LowerCaseName }}Repository: repos.{{ .CamelCaseName }}Repository(),
	}
}

func (u Create{{ .CamelCaseName }}UseCase) Execute(ctx context.Context, model models.{{ .CamelCaseName }}Model) (models.{{ .CamelCaseName }}Model, error) {
	return u.{{ .LowerCaseName }}Repository.Create(ctx, model)
}
{{ end }}
