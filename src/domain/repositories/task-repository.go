package repositories

import (
	"github.com/thunderjr/go-clean-api/src/domain/entities"
)

type TaskRepository interface {
	GetTask(taskId string) (entities.Task, error)
	SaveTask(task entities.Task) error
	UpdateTask(task entities.Task) error
}
