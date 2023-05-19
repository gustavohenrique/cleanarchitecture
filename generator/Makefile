repo_dir := $(CURDIR)/../
dist_dir := $(TMPDIR)/clean_architecture_generator/
download_dir := $(dist_dir)
resources_dir := $(CURDIR)/resources
binary := generator

clear:
	@rm -rf $(dist_dir) || echo

test: tests
tests: clear
	@DEBUG=true REPO_DIR=$(repo_dir) DIST_DIR=$(dist_dir) DOWNLOAD_DIR=$(download_dir) go test -v -cover ./src/...

run: clear
	@DEBUG=true REPO_DIR=$(repo_dir) DIST_DIR=$(dist_dir) DOWNLOAD_DIR=$(download_dir) go run main.go

build:
	rm $(binary)-* 2>/dev/null || echo
	@bash $(resources_dir)/build.sh $(binary)

update:
	@bash $(resources_dir)/update_repo.sh

deploy:
	@bash $(resources_dir)/deploy.sh $(binary)
