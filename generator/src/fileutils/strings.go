package fileutils

import (
	"path/filepath"
	"sort"
	"strings"
)

func isInvalidFile(filename string, extensions []string) bool {
	extension := filepath.Ext(filename)
	return strings.HasPrefix(filename, ".") || !sliceContains(extensions, extension)
}

func sliceContains(slice []string, term string) bool {
	sort.Strings(slice)
	for _, s := range slice {
		if strings.HasSuffix(s, term) {
			return true
		}
	}
	return false
}

func sliceContainsDir(slice []string, term string) bool {
	if sliceContains(slice, term) {
		return true
	}
	return sliceContains(slice, "/"+term+"/")
}

func contains(s string, terms []string) bool {
	for _, term := range terms {
		if strings.Contains(s, term) {
			return true
		}
	}
	return false
}

func hasEmpty(s ...string) bool {
	for _, i := range s {
		if trim(i) == "" {
			return true
		}
	}
	return false
}

func trim(s string) string {
	return strings.TrimSpace(s)
}
