package ports

import (
	"context"

	"{{ .ProjectName }}/src/domain/models"
)

{{ range .Models }}
type {{ .CamelCaseName }}UseCase interface {
	Execute(ctx context.Context, model models.{{ .CamelCaseName }}Model) (models.{{ .CamelCaseName }}Model, error)
}

type {{ .CamelCaseName }}Repository interface {
	Create(ctx context.Context, model models.{{ .CamelCaseName }}Model) (models.{{ .CamelCaseName }}Model, error)
	ReadOne(ctx context.Context, model models.{{ .CamelCaseName }}Model) (models.{{ .CamelCaseName }}Model, error)
	Update(ctx context.Context, model models.{{ .CamelCaseName }}Model) (models.{{ .CamelCaseName }}Model, error)
	Delete(ctx context.Context, model models.{{ .CamelCaseName }}Model) (models.{{ .CamelCaseName }}Model, error)
}
{{ end }}
