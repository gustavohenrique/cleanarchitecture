package fileutils_test

import (
	"testing"

	"generator/src/fileutils"
)

func TestGenerateNewFilesWithTheProjectName(t *testing.T) {
	sourceDir, distDir := getTemplateDirs()
	extensions := []string{".go", ".mod"}
	placeholders := map[string]string{
		"ProjectName": "mynewproject",
	}
	outputDir, err := fileutils.NewSed().From(sourceDir).To(distDir).Only(extensions).Replace(placeholders).Run()
	if err != nil || outputDir == "" {
		t.Errorf("Output dir=%s and error=%s", outputDir, err)
	}
}
