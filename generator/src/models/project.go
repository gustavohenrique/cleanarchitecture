package models

import (
	"fmt"
	"strings"
)

type Project struct {
	Name         string                 `json:"name"`
	Engine       string                 `json:"engine"`
	Placeholders map[string]interface{} `json:"placeholders"`
	Databases    []string               `json:"databases"`
	Servers      []string               `json:"servers"`
	Clients      []string               `json:"clients"`
	Sdks         []string               `json:"sdks"`
}

func NewProject(name, engine string) *Project {
	return &Project{
		Name:   name,
		Engine: engine,
		Placeholders: map[string]interface{}{
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

func (p *Project) IsValid() bool {
	hasEngine := p.Engine == GOLANG || p.Engine == QUASAR
	hasProjectName := len(p.GetName()) > 2
	hasAtLeatOneServer := len(p.Servers) > 0
	hasAtLeatOneDb := len(p.Databases) > 0
	return hasEngine && hasProjectName && hasAtLeatOneDb && hasAtLeatOneServer
}

func (p Project) String() string {
	return fmt.Sprintf("name=%s engine=%s servers=%s clientes=%s db=%s sdks=%s", p.GetName(), p.Engine, p.Servers, p.Clients, p.Databases, p.Sdks)
}
