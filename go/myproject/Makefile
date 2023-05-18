.PHONY: bin proto mocks run start install setup
go := $(shell which go)
scripts_dir := $(CURDIR)/scripts
configs_dir := $(CURDIR)/configs

default_config := $(configs_dir)/config.linux.yaml
ifeq ($(shell uname), Darwin)
default_config := $(configs_dir)/config.macos.yaml
endif
ifndef config
override config = $(default_config)
endif

include .makefiles/*

chmod:
	@chmod +x $(scripts_dir)/*

install: chmod
	$(go) install golang.org/x/tools/cmd/goimports@latest
	$(go) install github.com/golang/mock/mockgen@v1.6.0
	$(go) install honnef.co/go/tools/cmd/staticcheck@latest
	$(go) install github.com/swaggo/swag/cmd/swag@v1.8.4
	$(go) install github.com/axw/gocov/gocov@latest
	$(go) install github.com/AlekSi/gocov-xml@latest
	$(go) install github.com/gustavohenrique/coolconf/.../
	$(go) get github.com/golang/mock/gomock
	$(go) mod tidy

setup: install
	$(make) proto
	$(make) docs

start: run
run: chmod
	$(go) run cmd/webapp/main.go -config $(config)

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

swagger: docs
docs:
	@swag init -d ./src/infrastructure/servers/httpserver -g httpserver.go -o ./src/infrastructure/servers/httpserver/docs --parseDependency --parseInternal --parseDepth 5
	@echo
	@echo "Now open http://localhost:8001/docs/index.html"
