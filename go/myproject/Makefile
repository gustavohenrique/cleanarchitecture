.PHONY: bin proto mocks run start install setup
go := $(shell which go)
scripts_dir := $(CURDIR)/scripts
configs_dir := $(CURDIR)/configs
{{ if or .HasGrpcServer .HasGrpcWebServer .HasGrpcClient }}grpc := $(scripts_dir)/grpc.sh{{ end }}
{{ if or .HasGrpcWebServer }}grpcweb := $(scripts_dir)/grpcweb.sh{{ end }}
default_config := $(configs_dir)/config.linux.yaml
ifeq ($(shell uname), Darwin)
default_config := $(configs_dir)/config.macos.yaml
endif
ifndef config
override config = $(default_config)
endif

include .makefiles/*

setup: install {{ if .HasHttpServer}}docs{{ end}} {{ if or .HasGrpcServer .HasGrpcWebServer .HasGrpcClient }}install_grpc grpc{{ end }} {{ if .HasGrpcWebServer }}install_grpcweb grpcweb{{ end }} lint

chmod:
	@chmod +x $(scripts_dir)/*

install: chmod
	$(go) mod tidy
	$(go) install golang.org/x/tools/cmd/goimports@latest
	$(go) install github.com/golang/mock/mockgen@v1.6.0
	$(go) install honnef.co/go/tools/cmd/staticcheck@latest
	$(go) install github.com/axw/gocov/gocov@latest
	$(go) install github.com/AlekSi/gocov-xml@latest
	$(go) install github.com/gustavohenrique/coolconf/.../
	$(go) get github.com/golang/mock/gomock
	{{ if .HasHttpServer }}$(go) install github.com/swaggo/swag/cmd/swag@v1.8.4{{ end }}

start: run
run: chmod
	$(go) run cmd/app/main.go -config $(config)

test: tests
tests: chmod
	@clear
	@$(scripts_dir)/test.sh

mock: mocks
mocks: chmod
	@$(scripts_dir)/mockgen.sh

lint: mocks
	@goimports -w cmd src
	@staticcheck ./...
{{ if .HasHttpServer }}
swagger: docs
docs:
	@swag init -d ./src/infrastructure/servers/httpserver -g httpserver.go -o ./src/infrastructure/servers/httpserver/docs --parseDependency --parseInternal --parseDepth 5
	@echo
	@echo "API docs available in http://localhost:8001/docs/index.html"
{{ end }}
{{ if or .HasGrpcServer .HasGrpcWebServer .HasGrpcClient }}
install_grpc:
	@sh $(grpc) install
grpc:
	@sh $(grpc) generate
{{ end }}
{{ if .HasGrpcWebServer }}
install_grpcweb:
	@sh $(grpcweb) install
grpcweb:
	@sh $(grpcweb) generate
{{ end }}
