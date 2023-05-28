package models

{{ range .Models }}
type {{ .Name }}Model struct {
	Base
	{{ range .Fields }}
	{{ .GoName }} {{ .GoType }}
	{{ end }}
}
{{ end }}
