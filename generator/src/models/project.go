package models

import (
	"strings"
)

const GOLANG = "golang"
const QUASAR = "quasar"

type Project struct {
	Name         string            `json:"name"`
	Engine       string            `json:"engine"`
	Placeholders map[string]string `json:"placeholders"`
}

func NewProject(name, engine string) *Project {
	return &Project{
		Name:   name,
		Engine: engine,
		Placeholders: map[string]string{
			"ProjectName": name,
		},
	}
}

func (p *Project) GetName() string {
	return strings.ReplaceAll(strings.TrimSpace(p.Name), " ", "_")
}

func (p *Project) GetEngine() string {
	return p.Engine
}

func (p *Project) GetPlaceholders() map[string]string {
	placeholders := p.Placeholders
	if placeholders == nil {
		placeholders = map[string]string{
			"ProjectName": p.GetName(),
		}
	}
	return placeholders
}

func (p *Project) IsValid() bool {
	hasEngine := p.Engine == GOLANG || p.Engine == QUASAR
	hasProjectName := len(p.GetName()) > 2
	return hasEngine && hasProjectName
}
