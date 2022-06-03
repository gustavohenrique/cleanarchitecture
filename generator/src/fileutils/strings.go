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

func sliceContains(s []string, term string) bool {
	sort.Strings(s)
	i := sort.SearchStrings(s, term)
	return i < len(s) && s[i] == term
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
