package wire

import "{{ .ProjectName }}/src/domain/models"

{{ range .Models }}
type {{ .CamelCaseName }}HttpRequest struct {
	{{ range .Fields }}
	  {{ .GoName }} {{ .GoType }} `json:"{{ .GoName }}"`
	{{ end }}
}

func (wireIn *{{ .CamelCaseName }}HttpRequest) ToModel() models.{{ .CamelCaseName }}Model {
	return models.{{ .CamelCaseName }}Model{}
}

func (wireIn *{{ .CamelCaseName }}HttpRequest) Validate() error {
	return nil
}

type {{ .CamelCaseName }}HttpResponse struct {
	{{ range .Fields }}
	  {{ .GoName }} {{ .GoType }} `json:"{{ .GoName }}"`
	{{ end }}
}

func (wireOut {{ .CamelCaseName }}HttpResponse) Of(model models.{{ .CamelCaseName }}Model) {{ .CamelCaseName }}HttpResponse {
	{{ range .Fields }}
	wireOut.{{ .GoName }} = model.{{ .GoName }}
	{{ end }}
	return wireOut
}

type {{ .CamelCaseName }}Entity struct {
	{{ range .Fields }}
	  {{ .GoName }} {{ .GoType }} `json:"{{ .GoName }}"`
	{{ end }}
}

func (entity {{ .CamelCaseName }}Entity) ToModel() models.{{ .CamelCaseName }}Model {
	return models.{{ .CamelCaseName }}Model{
	{{ range .Fields }}
	{{ .GoName }}: entity.{{ .GoName }},
	{{ end }}
	}
}

func (entity {{ .CamelCaseName }}Entity) Of(m models.{{ .CamelCaseName }}Model) {{ .CamelCaseName }}Entity {
	return {{ .CamelCaseName }}Entity{
	{{ range .Fields }}
	{{ .GoName }}: m.{{ .GoName }},
	{{ end }}
	}
}
{{ end }}
