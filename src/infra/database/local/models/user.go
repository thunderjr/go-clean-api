package models

import (
	"github.com/thunderjr/go-clean-api/src/domain/entities"
)

type UserModel struct {
	BaseModel
	Name  string
	Tasks []*TaskModel `gorm:"foreignKey:UserID"`
}

func (u UserModel) ToUser() *entities.User {
	var tasks []entities.Task

	for _, task := range u.Tasks {
		tasks = append(tasks, *task.ToTask())
	}

	return &entities.User{
		ID:    u.ID.String(),
		Name:  u.Name,
		Tasks: tasks,
	}
}
