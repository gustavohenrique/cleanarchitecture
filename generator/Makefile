clear:
	@rm -rf $(CURDIR)/go_dist || echo
	@mkdir $(CURDIR)/go_dist
	@rm *.tar *.tar.gz || echo

test: tests
tests: clear
	@DEBUG=true SOURCE_DIR=$(CURDIR)/go_template/ DIST_DIR=$(CURDIR)/go_dist/ go test -v -cover ./src/...

run: clear
	@DEBUG=true SOURCE_DIR=$(CURDIR)/go_template/ DIST_DIR=$(CURDIR)/go_dist/ go run main.go
