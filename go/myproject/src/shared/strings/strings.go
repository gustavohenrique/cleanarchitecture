package strings

import (
	"sort"
	str "strings"
)

func HasEmpty(s ...string) bool {
	for _, i := range s {
		if Trim(i) == "" {
			return true
		}
	}
	return false
}

func Trim(s string) string {
	return str.TrimSpace(s)
}

func SliceContains(s []string, term string) bool {
	sort.Strings(s)
	i := sort.SearchStrings(s, term)
	return i < len(s) && s[i] == term
}
