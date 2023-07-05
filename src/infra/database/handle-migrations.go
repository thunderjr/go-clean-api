package infra_database

import (
	"sync"

	"github.com/thunderjr/go-clean-api/src/infra/database/local/models"
	"gorm.io/gorm"
)

var waitGroup sync.WaitGroup

func handleModelsChannel(db *gorm.DB, modelChannel chan interface{}, doneChannel chan error) {
	for model := range modelChannel {
		waitGroup.Add(1)

		go func(model interface{}) {
			err := db.AutoMigrate(model)
			waitGroup.Done()
			if err != nil {
				doneChannel <- err
			}
		}(model)
	}
}

func Migrate(db *gorm.DB) error {
	models := []any{
		&models.TaskModel{},
		&models.UserModel{},
	}

	modelChannel := make(chan interface{})
	doneChannel := make(chan error)
	go handleModelsChannel(db, modelChannel, doneChannel)

	for _, model := range models {
		modelChannel <- model
	}

	waitGroup.Wait()
	close(modelChannel)

	select {
	case err := <-doneChannel:
		return err
	default:
		return nil
	}
}
