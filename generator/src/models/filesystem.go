package models

import (
	"os"
	"path/filepath"
	"strings"

	"generator/src/fileutils/random"
)

const (
	REPO       = "repo"
	DIST       = "dist"
	EXTENSIONS = "extensions"
	SKIP       = "skip"
)

var engines = map[string]map[string]string{
	GOLANG: {
		REPO:       "go/myproject",
		DIST:       "go_projects",
		EXTENSIONS: ".go,.mod,.proto,.sh,.js",
		SKIP:       "node_modules,coverage,mocks,docs.go",
	},
	QUASAR: {
		REPO:       "js/quasar/myproject",
		DIST:       "quasar_projects",
		EXTENSIONS: ".js,.json",
		SKIP:       "node_modules,coverage",
	},
}

type Filesystem struct {
	repo     string
	dist     string
	download string
	engine   string
}

func getRepoByEngine(engine string) string {
	return engines[engine][REPO]
}

func getDistByEngine(engine string) string {
	return engines[engine][DIST]
}

func NewFilesystem(engine string) *Filesystem {
	repoDir, _ := filepath.Abs(filepath.Dir(os.Getenv("REPO_DIR")))
	repo := filepath.Join(repoDir, getRepoByEngine(engine))

	d, _ := filepath.Abs(filepath.Dir(os.Getenv("DIST_DIR")))
	distDir := filepath.Join(d, random.Strings(6))
	os.RemoveAll(distDir)
	os.MkdirAll(distDir, os.ModePerm)
	dist := filepath.Join(distDir, getDistByEngine(engine))

	downloadDir, _ := filepath.Abs(filepath.Dir(os.Getenv("DOWNLOAD_DIR")))
	os.MkdirAll(downloadDir, os.ModePerm)

	return &Filesystem{
		engine:   engine,
		repo:     repo,
		dist:     dist,
		download: downloadDir,
	}
}

func (d *Filesystem) GetExtensions() []string {
	return strings.Split(engines[d.engine][EXTENSIONS], ",")
}

func (d *Filesystem) GetSkipDirs(t *TemplateData) []string {
	skip := strings.Split(engines[d.engine][SKIP], ",")
	if t == nil {
		return skip
	}
	if len(t.Clients) == 0 {
		skip = append(skip, "clients")
	}
	if len(t.Sdks) == 0 {
		skip = append(skip, "sdk")
	}
	if !t.HasPostgres {
		skip = append(skip, "_"+POSTGRES, POSTGRES, POSTGRES+".sh")
	}
	if !t.HasSqlite {
		skip = append(skip, "_"+SQLITE, SQLITE, SQLITE+".sh")
	}
	if !t.HasDgraph {
		skip = append(skip, "_"+DGRAPH, DGRAPH, DGRAPH+".sh")
	}

	if !t.HasHttpServer && !t.HasGrpcServer {
		skip = append(skip, "assets")
	}
	if !t.HasHttpServer {
		skip = append(skip, HTTP+"server", "_"+HTTP, HTTP+"_", "web")
	}
	if !t.HasGrpcWebServer {
		skip = append(skip, GRPCWEB+"server", "_"+GRPCWEB, GRPCWEB+"_", GRPCWEB+".sh")
	}
	if !t.HasGrpcServer {
		skip = append(skip, GRPC+"server", "_"+GRPC+"_", GRPC+"_")
	}
	if !t.HasGrpcServer && !t.HasGrpcClient {
		skip = append(skip, GRPC+".sh")
	}
	if !t.HasGrpcServer && !t.HasGrpcWebServer && !t.HasGrpcClient {
		skip = append(skip, "proto")
	}
	if !t.HasNatsServer {
		skip = append(skip, NATS+"server", "_"+NATS, NATS+"_")
	}

	if !t.HasNatsClient {
		skip = append(skip, NATS+"client")
	}
	if !t.HasHttpClient {
		skip = append(skip, HTTP+"client")
	}
	if !t.HasGrpcClient {
		skip = append(skip, GRPC+"client")
	}

	if !t.HasJsHttpSdk {
		skip = append(skip, "js_rest")
	}
	if !t.HasGoGrpcSdk {
		skip = append(skip, "go_grpc")
	}
	return skip
}

func (d *Filesystem) GetRepo() string {
	return d.repo
}

func (d *Filesystem) GetDist() string {
	return d.dist
}

func (d *Filesystem) Dist(subfolder string) string {
	distDir := filepath.Join(d.GetDist(), subfolder)
	os.MkdirAll(distDir, os.ModePerm)
	return distDir
}

func (d *Filesystem) GetDownload() string {
	return d.download
}

func (d *Filesystem) Download(filename string) string {
	return filepath.Join(d.GetDownload(), filename)
}
