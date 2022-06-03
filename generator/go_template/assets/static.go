package assets

import (
	"embed"
	"io/fs"
	"net/http"
)

const staticFolder = "static"

//go:embed static
var staticFS embed.FS

type StaticFile struct {
	staticFS embed.FS
}

func NewStaticFile() StaticFile {
	return StaticFile{staticFS}
}

func (w StaticFile) GetFS() http.Handler {
	files, _ := fs.Sub(w.staticFS, staticFolder)
	httpFS := http.FS(files)
	return http.FileServer(httpFS)
}
