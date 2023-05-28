package repositories

import (
	"context"

	"{{ .ProjectName }}/src/domain/models"
	"{{ .ProjectName }}/src/domain/ports"
	"{{ .ProjectName }}/src/infrastructure/datastores/db"
	"{{ .ProjectName }}/src/wire"
)

{{ range .Models }}
type {{ .CamelCaseName }}SqliteRepository struct {
	store db.SqlDataStore
}

func New{{ .CamelCaseName }}SqliteRepository(store db.SqlDataStore) ports.{{ .CamelCaseName }}Repository {
	return &{{ .CamelCaseName }}SqliteRepository{store}
}

func (r {{ .CamelCaseName }}SqliteRepository) Create(ctx context.Context, model models.{{ .CamelCaseName }}Model) (models.{{ .CamelCaseName }}Model, error) {
	q := "INSERT INTO {{.SnakeCasePluralName}} ({{range .Fields}}{{.NameForSql}}, {{end}}) VALUES ({{range .Fields}}?, {{end}})"
	err := r.store.WithContext(ctx).Exec(q,{{ range .Fields }}
		model.{{ .NameForGo }},{{ end }}
	)
	return model, err
}

func (r {{ .CamelCaseName }}SqliteRepository) ReadOne(ctx context.Context, model models.{{ .CamelCaseName }}Model) (models.{{ .CamelCaseName }}Model, error) {
	var item wire.{{ .CamelCaseName }}Entity
	q := "SELECT {{range .Fields}}{{.NameForSql}}, {{end}} FROM {{.SnakeCasePluralName}} WHERE id=? LIMIT 1"
	err := r.store.WithContext(ctx).QueryAll(q, &item, model.ID)
	return item.ToModel(), err
}

func (r {{ .CamelCaseName }}SqliteRepository) Update(ctx context.Context, model models.{{ .CamelCaseName }}Model) (models.{{ .CamelCaseName }}Model, error) {
	q := "UPDATE {{.SnakeCasePluralName}} SET {{range .Fields}}{{.NameForSql}}=?, {{end}} WHERE id=$1"
	err := r.store.WithContext(ctx).Exec(q,{{ range .Fields }}
		model.{{ .NameForGo }},{{ end }}
	)
	return model, err
}

func (r {{ .CamelCaseName }}SqliteRepository) Delete(ctx context.Context, model models.{{ .CamelCaseName }}Model) (models.{{ .CamelCaseName }}Model, error) {
	q := "DELETE FROM {{.SnakeCasePluralName}} WHERE id=? LIMIT 1"
	err := r.store.WithContext(ctx).Exec(q, model.ID)
	return model, err
}
{{ end }}
