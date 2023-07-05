package local_repositories

import (
	uuid "github.com/satori/go.uuid"
	"github.com/thunderjr/go-clean-api/src/domain/entities"
	"github.com/thunderjr/go-clean-api/src/infra/database/local/models"
	"gorm.io/gorm"
)

type LocalTaskRepository struct {
	db *gorm.DB
}

func GetLocalTaskRepository(db *gorm.DB) LocalTaskRepository {
	return LocalTaskRepository{
		db,
	}
}

func (r *LocalTaskRepository) GetTask(taskId string) (entities.Task, error) {
	resultChan := make(chan entities.Task)
	errorChan := make(chan error)

	go func() {
		var task models.TaskModel
		if err := r.db.Where("id = ?", taskId).First(task).Error; err != nil {
			errorChan <- err
		}

		resultChan <- *task.ToTask()
	}()

	select {
	case task := <-resultChan:
		return task, nil
	case error := <-errorChan:
		return entities.Task{}, error
	}
}

func (r *LocalTaskRepository) SaveTask(task entities.Task) error {
	resultChan := make(chan error)

	go func() {
		modelTask := models.TaskModel{
			BaseModel: models.BaseModel{
				ID: uuid.NewV4(),
			},
			UserID: uuid.FromStringOrNil(task.UserID),
			Title:  task.Title,
		}

		if err := r.db.Create(&modelTask).Error; err != nil {
			resultChan <- err
		}

		resultChan <- nil
	}()

	return <-resultChan
}

func (r *LocalTaskRepository) UpdateTask(task entities.Task) error {
	resultChan := make(chan error)

	go func() {
		modelTask := models.TaskModel{
			BaseModel: models.BaseModel{
				ID: uuid.FromStringOrNil(task.ID),
			},
			UserID: uuid.FromStringOrNil(task.UserID),
			Title:  task.Title,
		}

		if err := r.db.Save(&modelTask).Error; err != nil {
			resultChan <- err
		}

		resultChan <- nil
	}()

	return <-resultChan
}
