package ports

type Repositories interface {
	{{ range .Models }}
	{{ .CamelCaseName }}Repository() {{ .CamelCaseName }}Repository
	{{ end }}
}
