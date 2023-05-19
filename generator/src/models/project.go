package models

import (
	"strings"
)

const (
	GOLANG     = "golang"
	QUASAR     = "quasar"
	HTTP       = "http"
	GRPC       = "grpc"
	GRPCWEB    = "grpcweb"
	NATS       = "nats"
	GO_GRPC    = "go_grpc"
	JS_HTTP    = "js_http"
	JS_GRCPWEB = "js_grpcweb"
)

type Project struct {
	Name         string            `json:"name"`
	Engine       string            `json:"engine"`
	Placeholders map[string]string `json:"placeholders"`
	Databases    []string          `json:"databases"`
	Servers      []string          `json:"servers"`
	Clients      []string          `json:"clients"`
	Sdks         []string          `json:"sdks"`
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
	hasAtLeatOneServer := len(p.Servers) > 0
	hasAtLeatOneDb := len(p.Databases) > 0
	return hasEngine && hasProjectName
}
