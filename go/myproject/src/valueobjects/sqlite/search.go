package sqlite

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strings"
)

/*
   Request example:
   {
       "placeholders": "($1 or $2) and $3 and $4",
       "filters": [
           {
               "placeholder": "$1",
               "field": "first_name",
               "operator": "contains",
               "term": "gus%"
           },
           {
               "placeholder": "2",
               "field": "last_name",
               "operator": "contains",
               "term": "henrique"
           },
           {
               "placeholder": "3",
               "field": "email",
               "operator": "eq",
               "term": "gustavo@gustavohenrique.net"
           },
           {
               "placeholder": "4",
               "field": "country",
               "operator": "any",
               "term": "brasil,argentina",
           }
       ],
       "pagination": {
           "page": 1,
           "per_page": 10
       },
       "sorting": {
           "field": "name",
           "order": "ASC"
       },
       "groupBy": {
		   "field": "name"
		}
   }
*/

const (
	ASC  = "asc"
	DESC = "desc"
)

const (
	EQUAL              = "eq"
	NOT_EQUAL          = "ne"
	GREATER_THAN       = "gt"
	LESS_THAN          = "lt"
	GREATER_THAN_EQUAL = "gte"
	LESS_THAN_EQUAL    = "lte"
	CONTAINS           = "contains"
	ANY                = "any"
)

var operators = map[string]string{
	EQUAL:              "=",
	NOT_EQUAL:          "<>",
	GREATER_THAN:       ">",
	LESS_THAN:          "<",
	GREATER_THAN_EQUAL: ">=",
	"ge":               ">=",
	LESS_THAN_EQUAL:    "<=",
	"le":               "<=",
	CONTAINS:           "LIKE",
	ANY:                "IN",
	"anyofterms":       "IN",
}

var allowedRawValues = []string{
	"CURRENT_TIMESTAMP",
	"NOW()",
}

type SearchRequest struct {
	Placeholders string         `json:"placeholders"`
	Pagination   Pagination     `json:"pagination"`
	Filters      []SearchFilter `json:"filters"`
	Sorting      Sorting        `json:"sorting"`
	Grouping     *Grouping      `json:"grouping"`
	Extra        interface{}    `json:"-"`
}

type SearchFilter struct {
	Placeholder string `json:"placeholder"`
	Field       string `json:"field"`
	Operator    string `json:"operator"`
	Term        string `json:"term"`
	Table       string `json:"table"`
}

type Sorting struct {
	Field string `json:"field"`
	Order string `json:"order"`
	Table string `json:"table"`
}

type Grouping struct {
	Field string `json:"field"`
	Table string `json:"table"`
}

type Pagination struct {
	Page    int32 `json:"page"`
	PerPage int32 `json:"per_page"`
	Total   int32 `json:"total"`
}

func (f Pagination) GetMaxPages(total int32) int32 {
	pages := int32(math.Ceil(float64(total) / float64(f.PerPage)))
	if pages == 0 {
		pages = 1
	}
	return pages
}

func Unmarshal(raw string) (SearchRequest, error) {
	var req SearchRequest
	err := json.Unmarshal([]byte(raw), &req)
	return req, err
}

func (f SearchRequest) String() string {
	sql := f.where()
	sql += f.groupBy()
	sql += f.orderBy()
	sql += f.limit()
	return sql
}

func (f SearchRequest) where() string {
	sql := "WHERE 1=1 "
	if f.Placeholders == "" || len(f.Filters) == 0 {
		return sql
	}
	if strings.Count(f.Placeholders, "$") != len(f.Filters) {
		return sql
	}
	sql += "AND"
	condition := f.Placeholders
	for _, filter := range f.Filters {
		term := removeDangerousWords(filter.Term)
		operator := operators[strings.ToLower(filter.Operator)]
		if operator == "" {
			operator = operators[EQUAL]
		}
		field := filter.Field
		if filter.Table != "" {
			field = filter.Table + "." + field
		}
		field = removeDangerousWords(field)
		expression := field + " " + operator + " " + quoteString(term)

		if operator == operators[ANY] {
			splited := strings.Split(term, ",")
			var terms []string
			for _, s := range splited {
				terms = append(terms, quoteString(strings.TrimSpace(s)))
			}
			term = strings.Join(terms, ", ")
			expression = field + " " + operator + " (" + term + ")"
		}
		placeholder := filter.Placeholder
		if !strings.HasPrefix(filter.Placeholder, "$") {
			placeholder = "$" + filter.Placeholder
		}
		condition = strings.Replace(condition, placeholder, expression, 1)
	}
	sql += " " + condition
	return sql
}

func (f SearchRequest) groupBy() string {
	if f.Grouping == nil {
		return ""
	}
	field := f.Grouping.Field
	if f.Grouping.Table != "" {
		field = f.Grouping.Table + "." + field
	}
	return " GROUP BY " + field
}

func (f SearchRequest) orderBy() string {
	sorting := f.Sorting
	if sorting.Field == "" || sorting.Order == "" {
		return ""
	}
	field := sorting.Field
	if sorting.Table != "" {
		field = sorting.Table + "." + field
	}
	return " ORDER BY " + field + " COLLATE NOCASE " + sorting.Order
}

func (f SearchRequest) getPerPageOrDefault() int32 {
	if f.Pagination.PerPage == 0 {
		return 10 // default
	}
	return f.Pagination.PerPage
}

func (f SearchRequest) limit() string {
	pagination := f.Pagination
	perPage := f.getPerPageOrDefault()
	page := int(math.Max(1, float64(pagination.Page)))
	offset := (page - 1) * int(perPage)
	return fmt.Sprintf(" LIMIT %d OFFSET %d", perPage, offset)
}

func quoteString(str string) string {
	// https://github.com/lib/pq/blob/master/conn.go#L1593
	if sliceContains(allowedRawValues, str) {
		return str
	}
	literal := strings.Replace(str, `'`, `''`, -1)
	if strings.Contains(literal, `\`) {
		literal = strings.Replace(literal, `\`, `\\`, -1)
		literal = ` E'` + literal + `'`
	} else {
		literal = `'` + literal + `'`
	}
	return literal
}

func sliceContains(s []string, term string) bool {
	sort.Strings(s)
	i := sort.SearchStrings(s, term)
	return i < len(s) && s[i] == term
}

func removeDangerousWords(str string) string {
	s := strings.ReplaceAll(str, "\xbf\x27", "")
	s = strings.ReplaceAll(s, "'", "")
	s = strings.ReplaceAll(s, "\"", "")
	s = strings.ReplaceAll(s, ";", "")
	return s
}
