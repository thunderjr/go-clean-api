package migrations

import (
	"log"
	"sync"

	"github.com/thunderjr/go-clean-api/src/infra/database/local/models"
	"gorm.io/gorm"
)

var waitGroup sync.WaitGroup

func handleModelsChannel(db *gorm.DB, modelChannel chan interface{}) {
	for model := range modelChannel {
		waitGroup.Add(1)

		go func(model interface{}) {
			defer waitGroup.Done()
			err := db.AutoMigrate(&model)
			if err != nil {
				log.Println(err)
			}
		}(model)
	}
}

func Migrate(db *gorm.DB) {
	models := []interface{}{
		models.TaskModel{},
		models.UserModel{},
	}

	modelChannel := make(chan interface{})
	go handleModelsChannel(db, modelChannel)

	for _, model := range models {
		modelChannel <- model
	}

	waitGroup.Wait()
}
