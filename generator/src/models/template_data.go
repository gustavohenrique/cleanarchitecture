package models

type TemplateData struct {
	ProjectName      string
	Databases        []string
	Servers          []string
	Clients          []string
	Sdks             []string
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
