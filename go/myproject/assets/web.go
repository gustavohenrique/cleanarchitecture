package assets

{{ if .HasHttpServer }}
import (
	"embed"
	"io/fs"
)

const folder = "web"

//go:embed web
var webFS embed.FS

type WebPage struct {
	webFS embed.FS
}

func NewWebPage() WebPage {
	return WebPage{webFS}
}

func (w WebPage) GetFS() fs.FS {
	files, _ := fs.Sub(w.webFS, folder)
	return files
}

func (w WebPage) Lookup(filename string) (string, error) {
	var content []byte
	err := fs.WalkDir(w.GetFS(), ".", func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if !d.IsDir() {
			if s == filename {
				b, err := w.webFS.ReadFile(folder + "/" + s)
				if err != nil {
					return err
				}
				content = b
			}
		}
		return nil
	})
	return string(content), err
}
{{ end }}
