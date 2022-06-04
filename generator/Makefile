source_dir := $(CURDIR)/../
dist_dir := /tmp/cleanarchitecturegenerator

clear:
	@rm -rf $(dist_dir) || echo
	@mkdir $(dist_dir)
	@rm *.tar *.tar.gz || echo

test: tests
tests: clear
	@DEBUG=true SOURCE_DIR=$(source_dir) DIST_DIR=$(dist_dir) go test -v -cover ./src/...

run: clear
	@DEBUG=true SOURCE_DIR=$(source_dir) DIST_DIR=$(dist_dir) go run main.go
