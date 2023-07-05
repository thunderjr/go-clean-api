package entities

import (
	"database/sql"
	"time"
)

type Task struct {
	ID        string
	Title     string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
