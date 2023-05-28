package models

import (
	"regexp"
	"strings"

	"generator/src/pluralize"
)

const (
	STRING = "string"
	BOOL   = "bool"
	INT    = "int"
	FLOAT  = "float"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

type TemplateModel struct {
	Name   string
	Fields []TemplateModelField
}

func (t TemplateModel) CamelCaseName() string {
	return strings.ReplaceAll(strings.TrimSpace(strings.Title(t.Name)), "Model", "")
}

func (t TemplateModel) LowerCaseName() string {
	return strings.ToLower(t.CamelCaseName())
}

func (t TemplateModel) SnakeCaseName() string {
	return toSnakeCase(t.Name)
}

func (t TemplateModel) SnakeCasePluralName() string {
	return toSnakeCase(pluralize.NewClient().Plural(t.Name))
}

type TemplateData struct {
	ProjectName      string
	Databases        []string
	Servers          []string
	Clients          []string
	Sdks             []string
	Models           []TemplateModel
	HasHttpServer    bool
	HasGrpcServer    bool
	HasGrpcWebServer bool
	HasNatsServer    bool
	HasHttpClient    bool
	HasGrpcClient    bool
	HasNatsClient    bool
	HasGoGrpcSdk     bool
	HasJsGrpcWebSdk  bool
	HasJsHttpSdk     bool
	HasPostgres      bool
	HasSqlite        bool
	HasDgraph        bool
}

func NewTemplateData() *TemplateData {
	return &TemplateData{}
}

func (t *TemplateData) Of(p Project) *TemplateData {
	hasPostgres := contains(p.Databases, POSTGRES)
	hasSqlite := contains(p.Databases, SQLITE)
	hasDgraph := contains(p.Databases, DGRAPH)
	hasHttpServer := contains(p.Servers, HTTP)
	hasGrpcServer := contains(p.Servers, GRPC)
	hasGrpcWebServer := contains(p.Servers, GRPCWEB)
	hasNatsServer := contains(p.Servers, NATS)
	hasHttpClient := contains(p.Clients, HTTP)
	hasGrpcClient := contains(p.Clients, GRPC)
	hasNatsClient := contains(p.Clients, NATS)
	hasGoGrpcSdk := contains(p.Sdks, GO_GRPC)
	hasJsGrpcWebSdk := contains(p.Sdks, JS_GRCPWEB)
	hasJsHttpSdk := contains(p.Sdks, JS_HTTP)

	t.ProjectName = p.GetName()
	t.Databases = p.Databases
	t.Servers = p.Servers
	t.Clients = p.Clients
	t.Sdks = p.Sdks

	t.HasPostgres = hasPostgres
	t.HasSqlite = hasSqlite
	t.HasDgraph = hasDgraph

	t.HasHttpServer = hasHttpServer || hasGrpcWebServer
	t.HasGrpcServer = hasGrpcServer
	t.HasGrpcWebServer = hasGrpcWebServer
	t.HasNatsServer = hasNatsServer

	t.HasHttpClient = hasHttpClient
	t.HasGrpcClient = hasGrpcClient
	t.HasNatsClient = hasNatsClient || hasNatsServer

	t.HasGoGrpcSdk = hasGoGrpcSdk && hasGrpcServer
	t.HasJsGrpcWebSdk = hasJsGrpcWebSdk || hasGrpcWebServer
	t.HasJsHttpSdk = hasJsHttpSdk && t.HasHttpServer

	models := []TemplateModel{}
	for _, projectModel := range p.Models {
		model := TemplateModel{}
		model.Name = projectModel.Name
		for _, projectModelField := range projectModel.Fields {
			model.Fields = append(model.Fields, TemplateModelField{
				Name: projectModelField.Name,
				Type: projectModelField.Type,
			})
		}
		models = append(models, model)
	}
	t.Models = models
	return t
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
