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

func (p *Project) GetTemplateData() map[string]interface{} {
	templateData := p.Placeholders
	if templateData == nil {
		templateData = map[string]interface{}{
			"ProjectName": p.GetName(),
		}
	}
	hasHttpServer := p.contains(p.Servers, HTTP)
	hasGrpcServer := p.contains(p.Servers, GRPC)
	hasGrpcWebServer := p.contains(p.Servers, GRPCWEB)
	hasNatsServer := p.contains(p.Servers, NATS)
	hasHttpClient := p.contains(p.Clients, HTTP)
	hasGrpcClient := p.contains(p.Clients, GRPC)
	hasNatsClient := p.contains(p.Clients, NATS)
	hasGoGrpcSdk := p.contains(p.Sdks, GO_GRPC)
	hasJsGrpcWebSdk := p.contains(p.Sdks, JS_GRCPWEB)
	hasJsHttpSdk := p.contains(p.Sdks, JS_HTTP)

	templateData["HasHttpServer"] = hasHttpServer
	templateData["HasGrpcServer"] = hasGrpcServer
	templateData["HasGrpcWebServer"] = hasGrpcWebServer && hasHttpServer
	templateData["HasNatsServer"] = hasNatsServer
	templateData["HasHttpClient"] = hasHttpClient
	templateData["HasGrpcClient"] = hasGrpcClient
	templateData["HasNatsClient"] = hasNatsClient && hasNatsServer
	templateData["HasGoGrpcSdk"] = hasGoGrpcSdk && hasGrpcServer
	templateData["HasJsGrpcWebSdk"] = hasJsGrpcWebSdk && hasGrpcWebServer
	templateData["HasJsHttpSdk"] = hasJsHttpSdk && hasHttpServer
	return templateData
}

func (p *Project) IsValid() bool {
	hasEngine := p.Engine == GOLANG || p.Engine == QUASAR
	hasProjectName := len(p.GetName()) > 2
	hasAtLeatOneServer := len(p.Servers) > 0
	hasAtLeatOneDb := len(p.Databases) > 0
	return hasEngine && hasProjectName && hasAtLeatOneDb && hasAtLeatOneServer
}

func (p *Project) contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
