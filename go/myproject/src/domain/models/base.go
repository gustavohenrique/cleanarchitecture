package models

import "time"

type Base struct {
	ID        string
	CreatedAt *time.Time
}
