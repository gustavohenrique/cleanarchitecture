package assets

import (
	"embed"
	"io/fs"
	"net/http"
)

const folder = "html"

//go:embed html
var assetsFS embed.FS

type Assets struct {
	assetFS embed.FS
}

func New() *Assets {
	return &Assets{assetsFS}
}

func (w *Assets) GetFS() http.Handler {
	files, _ := fs.Sub(w.assetFS, folder)
	httpFS := http.FS(files)
	return http.FileServer(httpFS)
}
