package dgraph

import (
	"fmt"

	"{{ .ProjectName }}/src/shared/strings"
)

const OperatorAny = "anyofterms"
const OperatorAll = "allofterms"
const OperatorExact = "eq"

var validOperators = []string{OperatorAny, OperatorAll, OperatorExact}

type SearchRequest struct {
	SortFieldAsc  string       `json:"asc"`
	SortFieldDesc string       `json:"desc"`
	PerPage       int32        `json:"per_page"`
	Page          int32        `json:"page"`
	Total         int32        `json:"total"`
	Filter        SearchFilter `json:"filter"`
}

func (s SearchRequest) String() string {
	return "NOT IMPLEMENTED"
}

type SearchFilter struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Term     string `json:"term"`
}

func (s SearchRequest) ParsePagination() string {
	page := s.Page - 1
	perPage := s.PerPage
	if page < 0 || perPage <= 0 {
		page = 0
		perPage = 10
	}
	return fmt.Sprintf("first: %d, offset: %d", perPage, page)
}

func (s SearchFilter) ValidateOperator() error {
	if strings.SliceContains(validOperators, s.Operator) {
		return nil
	}
	return fmt.Errorf("Unknown operator %s", s.Operator)
}

func (s SearchRequest) ParseFilter() (string, error) {
	filter := s.Filter
	if strings.HasEmpty(filter.Operator, filter.Field, filter.Term) {
		return "", nil
	}
	if err := filter.ValidateOperator(); err != nil {
		return "", err
	}
	return fmt.Sprintf(`@filter(%s(%s, "%s"))`, filter.Operator, filter.Field, filter.Term), nil
}

func (s SearchRequest) ParseSorting() string {
	asc := strings.Trim(s.SortFieldAsc)
	if asc == "" {
		asc = "created_at"
	}
	desc := strings.Trim(s.SortFieldDesc)
	if desc == "" {
		desc = "created_at"
	}
	if asc == desc {
		return "orderasc: " + asc
	}
	return "orderasc: " + asc + ", orderdesc: " + desc
}
