.PHONY: releases proto mocks

go := $(shell which go)
scripts_dir := $(CURDIR)/scripts
pgoptions := "-q -X -v VERBOSITY=terse -v ON_ERROR_STOP=1 -A -t -w"
postgres_user := root
postgres_password := $(postgres_user)
postgres_db := $(postgres_user)
postgres_url := postgres://$(postgres_user):$(postgres_password)@127.0.0.1/$(postgres_db)?sslmode=disable
watch_files := (.go$$)|(.html$$)
proto_source_dir := assets/static/proto
go_proto_out_dir := src/proto
grpc_web_proto_out_dir := sdk/js_grpcweb
grpc_web_version := 1.3.1

setup: install
install:
	$(go) install google.golang.org/protobuf/cmd/protoc-gen-go
	$(go) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	$(go) install golang.org/x/tools/cmd/goimports@latest
	$(go) install github.com/cespare/reflex@latest
	$(go) install github.com/golang/mock/mockgen@v1.6.0
	$(go) install coolconf/.../
	$(go) mod tidy
	@echo
	@echo "Download protoc-gen-grpc and put it on PATH:"
	@echo "https://github.com/grpc/grpc-web/releases/download/$(grpc_web_version)/protoc-gen-grpc-web-$(grpc_web_version)-linux-x86_64"
	@echo "Do not forget to install manually the protobuf-all binary found in:"
	@echo "https://github.com/protocolbuffers/protobuf/releases"
	@echo "Unzip protobuf-all-*.zip and run ./configure && make && make install"

run:
	$(go) run cmd/server/main.go -config config.local.yaml

watch:
	reflex -r "$(watch_files)" -s -- make run

test: tests
tests:
	@clear
	@chmod +x $(scripts_dir)/*
	@$(scripts_dir)/test.sh

mock: mocks
mocks:
	@./$(scripts_dir)/mocks.sh

lint:
	@goimports -w cmd src

build:
	@bash -c "DEV_MODE=1 '$(scripts_dir)/build.sh'"

releases: lint
	@bash -c "RELEASE_MODE=1 '$(scripts_dir)/build.sh'"

migrate:
ifeq ($(strip $(db)), postgres)
	@$(MAKE) migrate_postgres
else
	@$(MAKE) migrate_sqlite
endif

migrate_sqlite:
	@#bash -c "DB_ENGINE=sqlite DATABASE_URL=$(CURDIR)/database.db MIGRATIONS_DIR=$(CURDIR)/migrations/sqlite '$(scripts_dir)/sql.sh' menu"
	@bash -c "DB_ENGINE=sqlite DATABASE_URL=$(CURDIR)/database.db MIGRATIONS_DIR=$(CURDIR)/migrations/sqlite '$(scripts_dir)/sql.sh' apply"

sqlite:
	@rm database.db 2>/dev/null || echo
	@sqlite3 database.db < ./migrations/sqlite/schema.sql
	@$(MAKE) migrate db=sqlite

migrate_postgres:
	@bash -c "DB_ENGINE=postgres DATABASE_URL=$(postgres_url) MIGRATIONS_DIR=$(CURDIR)/migrations/postgres '$(scripts_dir)/sql.sh' menu"

postgres: drop_postgres
	@docker run -d --name postgres \
		-m 2G \
		-e TZ=America/Sao_Paulo \
		-e POSTGRES_USER=$(postgres_user) \
		-e POSTGRES_PASSWORD=$(postgres_password) \
		-e POSTGRES_DB=$(postgres_db) \
		-p 5432:5432 \
		gustavohenrique/postgres:13-alpine \
	&& echo "Waiting for 3s..." \
	&& sleep 3

drop_postgres:
	@docker rm -f postgres 2>/dev/null || exit 0

psql:
	@docker exec -it postgres sh -c "PGPASSWORD=$(postgres_password) psql -U $(postgres_user) -d $(postgres_db) $(pgoptions)"

proto:
	@echo "Generating Go stubs..."
	@mkdir -p $(go_proto_out_dir) || echo
	@protoc -I $(proto_source_dir) $(proto_source_dir)/*.proto --proto_path=$(go_proto_out_dir) --go_out=$(go_proto_out_dir) --go-grpc_out=$(go_proto_out_dir)
	@echo "Generating JS stubs for gRPC Web..."
	@protoc -I $(proto_source_dir) $(proto_source_dir)/*.proto --js_out=import_style=commonjs,binary:$(grpc_web_proto_out_dir) --grpc-web_out=import_style=commonjs,mode=grpcweb:$(grpc_web_proto_out_dir)
	@echo "Copying Go stubs to SDK Go dir..."
	@cp -r $(go_proto_out_dir) sdk/go_grpc/
	@echo "Done."
