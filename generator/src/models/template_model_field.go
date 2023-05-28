package models

type TemplateModelField struct {
	Name string
	Type string
}

func (t TemplateModelField) GoName() string {
	return t.Name
}

func (t TemplateModelField) GoType() string {
	if t.Type == FLOAT {
		return "float32"
	}
	return t.Type
}

func (t TemplateModelField) NameForGo() string {
	return t.Name
}

func (t TemplateModelField) TypeForGo() string {
	if t.Type == FLOAT {
		return "float32"
	}
	return t.Type
}

func (t TemplateModelField) NameForProtobuf() string {
	return toSnakeCase(t.Name)
}

func (t TemplateModelField) TypeForProtobuf() string {
	if t.Type == INT {
		return "int32"
	}
	return t.Type
}

func (t TemplateModelField) NameForSql() string {
	return toSnakeCase(t.Name)
}

func (t TemplateModelField) TypeForSql() string {
	switch t.Type {
	case INT:
		return "INTEGER"
	case STRING:
		return "VARCHAR"
	case BOOL:
		return "BOOLEAN"
	case FLOAT:
		return "DECIMAL"
	default:
		return "VARCHAR"
	}
}
