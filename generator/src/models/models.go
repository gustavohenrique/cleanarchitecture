package models

import (
	"os"
	"path/filepath"
)

const GOLANG = "golang"
const JAVASCRIPT = "javascript"

type Project struct {
	Language     string            `json:"language"`
	Placeholders map[string]string `json:"placeholders"`
}

func (p *Project) GetFileExtensionsToBeReplaced() []string {
	if p.Language == JAVASCRIPT {
		return []string{".js", ".vue", ".json"}
	}
	return []string{".go", ".mod"}
}

func (p *Project) GetSkipDirs() []string {
	return []string{"node_modules", "coverage"}
}

func (p *Project) GetPlaceholders() map[string]string {
	return p.Placeholders
}

func (p *Project) IsValid() bool {
	hasLanguage := p.Language == GOLANG || p.Language == JAVASCRIPT
	return hasLanguage
}

func (p *Project) GetTemplatesDirs() (string, string) {
	sourceDir, _ := filepath.Abs(filepath.Dir(os.Getenv("SOURCE_DIR")))
	distDir, _ := filepath.Abs(filepath.Dir(os.Getenv("DIST_DIR")))
	source := "go/myproject"
	dist := "mygoproject"
	if p.Language == JAVASCRIPT {
		source = "js/quasar/myproject"
		dist = "myjsproject"
	}
	return filepath.Join(sourceDir, source), filepath.Join(distDir, dist)
}
