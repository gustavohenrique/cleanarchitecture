package models

type TodoModel struct {
	Base
	Title  string
	IsDone bool
}
