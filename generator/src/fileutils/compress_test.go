package fileutils_test

import (
	"path/filepath"
	"testing"

	"generator/src/fileutils"
)

func TestCompressTarGzDir(t *testing.T) {
	sourceDir, _ := getTemplateDirs()
	file, err := fileutils.NewCompress().Target(sourceDir).Exclude([]string{"node_modules", "coverage"}).Run()
	expected := filepath.Join(filepath.Dir(sourceDir), "go_template.tar.gz")
	if err != nil || file != expected {
		t.Errorf("Error: %s. Expected %s but got %s", err, expected, file)
	}
}
