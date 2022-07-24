.PHONY: bin proto mocks
go := $(shell which go)
scripts_dir := $(CURDIR)/scripts

include .makefiles/*

chmod:
	@chmod +x $(scripts_dir)/*

setup: install
install: chmod
	$(go) install golang.org/x/tools/cmd/goimports@latest
	$(go) install github.com/golang/mock/mockgen@v1.6.0
	$(go) install honnef.co/go/tools/cmd/staticcheck@latest
	$(go) install github.com/axw/gocov/gocov@latest
	$(go) install github.com/AlekSi/gocov-xml@latest
	$(go) install coolconf/.../
	$(go) mod tidy

run: chmod
	$(go) run cmd/server/main.go -config config.local.yaml

test: tests
tests: chmod
	@clear
	@$(scripts_dir)/test.sh

mock: mocks
mocks: chmod
	@./$(scripts_dir)/mocks.sh

lint:
	@goimports -w cmd src
	@staticcheck ./...

