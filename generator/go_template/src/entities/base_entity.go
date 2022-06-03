package entities

import "time"

type Base struct {
	ID         string     `json:"id" db:"id" dgraph:"uid,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty" db:"created_at" dgraph:"created_at,omitempty"`
	DgraphType string     `json:"-" db:"-" dgraph:"dgraph.type"`
}
