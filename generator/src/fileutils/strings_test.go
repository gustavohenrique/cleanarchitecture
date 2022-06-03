package fileutils

import (
	"testing"
)

func TestStringContainsAnyTerm(t *testing.T) {
	s := "/home/gustavo/Workspace/clean-architecture/generator/go_template/src/valueobjects"
	terms := []string{"node_modules", "coverage"}
	has := contains(s, terms)
	if has {
		t.Error("Expected false but got", has)
	}
}
