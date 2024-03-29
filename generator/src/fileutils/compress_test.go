package fileutils_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"generator/src/fileutils"
	"generator/src/models"
)

func TestCompressTarGzDir(t *testing.T) {
	name := "mytestproject"
	engine := models.GOLANG
	filesystem := models.NewFilesystem(engine)
	downloadDir := filesystem.GetDownload()
	project := models.NewProject(name, engine)
	file, err := fileutils.
		NewCompress().
		Input(filesystem.GetRepo()).
		Output(downloadDir).
		Name(name).
		Exclude(filesystem.GetSkipDirs(project)).
		Run()
	expected := filepath.Join(downloadDir, fmt.Sprintf("%s.tar.gz", name))
	if err != nil || file != expected {
		t.Errorf("Error: %s. Expected %s but got %s", err, expected, file)
	}
}
