package fileutils_test

import (
	"testing"

	"generator/src/fileutils"
	"generator/src/models"
)

func TestGenerateNewFilesWithTheProjectName(t *testing.T) {
	project := models.NewProject("mynewproject", models.GOLANG)
	filesystem := models.NewFilesystem(project.GetEngine())
	outputDir, err := fileutils.
		NewSed().
		From(filesystem.GetRepo()).
		To(filesystem.Dist(project.GetName())).
		Exclude([]string{"docs"}).
		Only(filesystem.GetExtensions()).
		Replace(models.NewTemplateData().Of(*project)).
		Run()
	if err != nil || outputDir == "" {
		t.Errorf("Output dir=%s and error=%s", outputDir, err)
	}
}
