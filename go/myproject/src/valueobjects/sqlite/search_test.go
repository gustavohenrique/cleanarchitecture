package sqlite_test

import (
	"testing"

	"myproject/src/valueobjects/sqlite"
)

var req sqlite.SearchRequest

func beforeEach() {
	req = sqlite.SearchRequest{}
	req.Placeholders = "$1 or $2"
	term1 := sqlite.SearchFilter{
		Placeholder: "1",
		Field:       "name",
		Operator:    "eq",
		Term:        "gustavo",
	}
	term2 := sqlite.SearchFilter{
		Placeholder: "2",
		Field:       "age",
		Operator:    "gte",
		Term:        "20",
	}
	req.Filters = []sqlite.SearchFilter{term1, term2}
	req.Pagination = sqlite.Pagination{
		PerPage: 10,
		Page:    1,
	}
	req.Sorting = sqlite.Sorting{
		Field: "id",
		Order: "DESC",
	}
}

func TestContainsShouldUseUnaccent(t *testing.T) {
	req = sqlite.SearchRequest{}
	req.Filters = []sqlite.SearchFilter{
		sqlite.SearchFilter{
			Placeholder: "1",
			Term:        "%questões%",
			Field:       "name",
			Operator:    "contains",
		},
		sqlite.SearchFilter{
			Placeholder: "2",
			Term:        "36",
			Field:       "age",
			Operator:    "eq",
		},
	}
	req.Placeholders = "$1 AND $2"
	sql := req.String()
	expected := "WHERE 1=1 AND name LIKE '%questões%' AND age = '36' LIMIT 10 OFFSET 0"
	if expected != sql {
		t.Errorf("EXPECTED: %s | GOT: %s", expected, sql)
	}
}

func TestReturnParametersAccordingOfRequest(t *testing.T) {
	beforeEach()
	sql := req.String()
	expected := "WHERE 1=1 AND name = 'gustavo' or age >= '20' ORDER BY id COLLATE NOCASE DESC LIMIT 10 OFFSET 0"
	if expected != sql {
		t.Errorf("EXPECTED: %s | GOT: %s", expected, sql)
	}
}

func TestDefaultSorting(t *testing.T) {
	beforeEach()
	req.Sorting = sqlite.Sorting{}
	sql := req.String()
	if "WHERE 1=1 AND name = 'gustavo' or age >= '20' LIMIT 10 OFFSET 0" != sql {
		t.Errorf("Failed")
	}
}

func TestDefaultPagination(t *testing.T) {
	beforeEach()
	req.Pagination = sqlite.Pagination{}
	sql := req.String()
	if "WHERE 1=1 AND name = 'gustavo' or age >= '20' ORDER BY id COLLATE NOCASE DESC LIMIT 10 OFFSET 0" != sql {
		t.Errorf("Failed")
	}
}

func TestUsingAnyInArrays(t *testing.T) {
	req = sqlite.SearchRequest{}
	req.Placeholders = "$1 AND $2"
	term1 := sqlite.SearchFilter{
		Placeholder: "1",
		Field:       "name",
		Operator:    sqlite.EQUAL,
		Term:        "gustavo",
	}
	term2 := sqlite.SearchFilter{
		Placeholder: "2",
		Field:       "country",
		Operator:    sqlite.ANY,
		Term:        "brasil, argentina",
	}
	req.Filters = []sqlite.SearchFilter{term1, term2}
	sql := req.String()
	expected := "WHERE 1=1 AND name = 'gustavo' AND country IN ('brasil', 'argentina') LIMIT 10 OFFSET 0"
	if expected != sql {
		t.Errorf("Failed. Expected %s but got %s", expected, sql)
	}
}

func TestIgnoreTermWhenItIsEmpty(t *testing.T) {
	beforeEach()
	req.Filters = []sqlite.SearchFilter{}
	sql := req.String()
	if "WHERE 1=1  ORDER BY id COLLATE NOCASE DESC LIMIT 10 OFFSET 0" != sql {
		t.Errorf("Failed")
	}
}

func TestIgnoreRequestWhenTotalTermsMismatchTotalParams(t *testing.T) {
	beforeEach()
	req.Placeholders = "$1 or $2 and $3"
	sql := req.String()
	if "WHERE 1=1  ORDER BY id COLLATE NOCASE DESC LIMIT 10 OFFSET 0" != sql {
		t.Errorf("Failed")
	}
}

func TestSQLInjection(t *testing.T) {
	beforeEach()
	req.Filters[0] = sqlite.SearchFilter{
		Placeholder: "1",
		Field:       "name",
		Operator:    sqlite.EQUAL,
		Term:        "\xbf\x27 or (select amount from orders limit 1) as created_at or name='",
	}
	sql := req.String()
	expected := "WHERE 1=1 AND name = ' or (select amount from orders limit 1) as created_at or name=' or age >= '20' ORDER BY id COLLATE NOCASE DESC LIMIT 10 OFFSET 0"
	if expected != sql {
		t.Errorf("EXPECTED: %s | GOT: %s", expected, sql)
	}
}
