package templaterender

import (
	"bytes"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Renderer struct {
	templates *template.Template
}

func NewRenderer(templates *template.Template) *Renderer {
	return &Renderer{templates}
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Parse(content string, c interface{}) string {
	var b bytes.Buffer
	t, err := template.New("").Parse(content)
	if err != nil {
		return "[ERROR] Failed to parse template content"
	}
	err = t.Execute(&b, c)
	if err != nil {
		return "[ERROR] Failed to execute template content"
	}
	return b.String()
}
