package models

type Project struct {
	Language     string            `json:"language"`
	Placeholders map[string]string `json:"placeholders"`
}

func (p *Project) GetFileExtensionsToBeReplaced() []string {
	if p.Language == "javascript" {
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
