package customjson_test

import (
	"testing"

	"{{ .ProjectName }}/src/infrastructure/datastores/dgraph/customjson"
)

func TestMarshalJsonTag(t *testing.T) {
	type S struct {
		ID string `json:"id" sometag:"uid"`
	}
	data := S{ID: "xpto"}
	res := customjson.MarshalWithCustomTag("json", data)
	expected := `{"id":"xpto"}`
	if expected != res {
		t.Fatalf("Expected %s and got %s", expected, res)
	}
}

func TestMarshalSimpleStructWithCustomTag(t *testing.T) {
	type S struct {
		ID string `json:"id" dgraph:"uid"`
	}
	data := S{ID: "xpto"}
	res := customjson.MarshalWithCustomTag("dgraph", data)
	expected := `{"uid":"xpto"}`
	if expected != res {
		t.Fatalf("Expected %s and got %s", expected, res)
	}
}

func TestMarshalSimpleStructIgnoringSomeField(t *testing.T) {
	type S struct {
		ID   string `json:"id" dgraph:"uid"`
		Name string `json:"name" dgraph:"-"`
	}
	data := S{
		ID:   "xpto",
		Name: "someone",
	}
	res := customjson.MarshalWithCustomTag("dgraph", data)
	expected := `{"uid":"xpto"}`
	if expected != res {
		t.Fatalf("Expected %s and got %s", expected, res)
	}
}

func TestMarshalSimpleStructIgnoringOmitEmpty(t *testing.T) {
	type S struct {
		ID   string `dgraph:"uid,omitempty"`
		Name string `dgraph:"full_name"`
	}
	data := S{
		Name: "someone",
	}
	res := customjson.MarshalWithCustomTag("dgraph", data)
	expected := `{"full_name":"someone"}`
	if expected != res {
		t.Fatalf("Expected %s and got %s", expected, res)
	}
}

func TestMarshalComplexStruct(t *testing.T) {
	type Employee struct {
		Name string `json:"name" dgraph:"full_name"`
	}
	type Company struct {
		ID       string   `json:"id" dgraph:"uid"`
		Employee Employee `json:"employee" dgraph:"person"`
	}
	company := Company{
		ID:       "xpto",
		Employee: Employee{Name: "gustavo"},
	}
	res := customjson.MarshalWithCustomTag("dgraph", company)
	expected := `{"uid":"xpto","person":{"full_name":"gustavo"}}`
	if expected != res {
		t.Fatalf("Expected %s and got %s", expected, res)
	}
}

func TestMarshalStructWithPointer(t *testing.T) {
	type Employee struct {
		Name string `dgraph:"full_name"`
	}
	type Company struct {
		ID       string    `dgraph:"uid"`
		Employee *Employee `dgraph:"person"`
	}
	company := Company{
		ID: "xpto",
		Employee: &Employee{
			Name: "gustavo",
		},
	}
	res := customjson.MarshalWithCustomTag("dgraph", company)
	expected := `{"uid":"xpto","person":{"full_name":"gustavo"}}`
	if expected != res {
		t.Fatalf("Expected %s and got %s", expected, res)
	}
}

func TestMarshalSlice(t *testing.T) {
	type A struct {
		ID string `dgraph:"uid"`
	}
	type B struct {
		Name string `dgraph:"name"`
	}
	a := A{ID: "0x1"}
	b := B{Name: "someone"}
	data := []interface{}{a, b}
	res := customjson.MarshalWithCustomTag("dgraph", data)
	expected := `[{"uid":"0x1"},{"name":"someone"}]`
	if expected != res {
		t.Fatalf("Expected %s and got %s", expected, res)
	}
}

func TestUnmarshalFromString(t *testing.T) {
	type Employee struct {
		Name string `dgraph:"full_name"`
	}
	type Company struct {
		ID       string    `dgraph:"uid"`
		Employee *Employee `dgraph:"person"`
	}
	var company Company
	jsonStr := `{
        "uid": "0x1",
		"person": {
		    "full_name": "someone"
		}
    }`
	customjson.UnmarshalWithCustomTag("dgraph", jsonStr, &company)
	if company.Employee.Name != "someone" {
		t.Fatalf("Expected 'someone' and got '%s'", company.Employee.Name)
	}
}
