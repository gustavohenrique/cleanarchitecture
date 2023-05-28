package repositories

import (
	"{{ .ProjectName }}/src/domain/ports"
	"{{ .ProjectName }}/src/infrastructure/datastores"
)

type RepositoriesContainer struct {
{{ range .Models }}
	{{ .LowerCaseName }}Repository ports.{{ .CamelCaseName }}Repository
{{ end }}
}

func New(datastores datastores.Stores) ports.Repositories {
	repos := &RepositoriesContainer{}
{{ range .Models }}
	{{ if $.HasPostgres }}
	repos.{{ .LowerCaseName }}Repository = New{{ .CamelCaseName }}PostgresRepository(datastores.Postgres())
	{{ else }}
		{{ if $.HasSqlite }}
	repos.{{ .LowerCaseName }}Repository = New{{ .CamelCaseName }}SqliteRepository(datastores.Sqlite())
		{{ else }}
			{{ if $.HasDgraph }}
	repos.{{ .LowerCaseName }}Repository = New{{ .CamelCaseName }}DgraphRepository(datastores.Dgraph())
			{{ end }}
		{{ end }}
	{{ end }}
{{ end }}
	return repos
}
{{ range .Models }}
func (repos *RepositoriesContainer) {{ .CamelCaseName }}Repository() ports.{{ .CamelCaseName }}Repository {
	return repos.{{ .LowerCaseName }}Repository
}
{{ end }}
