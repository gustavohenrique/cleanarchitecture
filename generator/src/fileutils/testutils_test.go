package fileutils_test

import (
	"os"
	"path/filepath"
)

func getTemplateDirs() (string, string) {
	sourceDir, _ := filepath.Abs(filepath.Dir(os.Getenv("SOURCE_DIR")))
	distDir, _ := filepath.Abs(filepath.Dir(os.Getenv("DIST_DIR")))
	return sourceDir, distDir
}
