package models

import (
	"database/sql"

	uuid "github.com/satori/go.uuid"
	"github.com/thunderjr/go-clean-api/src/domain/entities"
)

type TaskModel struct {
	BaseModel
	Title  string
	UserID uuid.UUID
}

func (t *TaskModel) ToTask() *entities.Task {
	return &entities.Task{
		ID:        t.ID.String(),
		Title:     t.Title,
		UserID:    t.UserID.String(),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		DeletedAt: sql.NullTime(t.DeletedAt),
	}
}
