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
	GOLANG: map[string]string{
		REPO:       "go/myproject",
		DIST:       "go_projects",
		EXTENSIONS: ".go,.mod,.proto,.sh,.js",
		SKIP:       "node_modules,coverage,mocks",
	},
	QUASAR: map[string]string{
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

func (d *Filesystem) GetSkipDirs() []string {
	return strings.Split(engines[d.engine][SKIP], ",")
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
