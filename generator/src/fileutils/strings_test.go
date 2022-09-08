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

func TestStringSliceContains(t *testing.T) {
	term := "mocks"
	slice := []string{
		"/tmp/mocks",
		"/somedir",
	}
	contains := sliceContains(slice, term)
	if !contains {
		t.Errorf("Expected to catch /tmp/mocks")
	}
	slice = []string{"/tmp/mocks.sh", "/anything"}
	contains = sliceContains(slice, term)
	if contains {
		t.Errorf("Expected to ignore /tmp/mocks.sh")
	}
}
