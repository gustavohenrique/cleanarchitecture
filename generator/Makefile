repo_dir := $(CURDIR)/../
dist_dir := /tmp/clean_architecture_generator/
download_dir := $(dist_dir)

clear:
	@rm -rf $(dist_dir) || echo

test: tests
tests: clear
	@DEBUG=true REPO_DIR=$(repo_dir) DIST_DIR=$(dist_dir) DOWNLOAD_DIR=$(download_dir) go test -v -cover ./src/...

run: clear
	@DEBUG=true REPO_DIR=$(repo_dir) DIST_DIR=$(dist_dir) DOWNLOAD_DIR=$(download_dir) go run main.go
