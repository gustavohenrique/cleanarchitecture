.PHONY: bin proto mocks
go := $(shell which go)
scripts_dir := $(CURDIR)/scripts

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

run: chmod
	$(go) run cmd/server/main.go -config config.local.yaml

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

docs:
	@swag init -d ./src/application/httpserver -g httpserver.go -o src/application/httpserver/docs --parseDependency --parseInternal --parseDepth 5
	@echo
	@echo "Now open http://localhost:8001/docs/index.html"
