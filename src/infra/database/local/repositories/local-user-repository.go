package local_repositories

import (
	uuid "github.com/satori/go.uuid"
	"github.com/thunderjr/go-clean-api/src/domain/entities"
	"github.com/thunderjr/go-clean-api/src/domain/repositories"
	"github.com/thunderjr/go-clean-api/src/infra/database/local/models"
	"gorm.io/gorm"
)

type LocalUserRepository struct {
	db *gorm.DB
}

func GetLocalUserRepository(db *gorm.DB) repositories.UserRepository {
	return &LocalUserRepository{
		db,
	}
}

func (r *LocalUserRepository) GetUser(userId string) (entities.User, error) {
	resultChan := make(chan entities.User)
	errorChan := make(chan error)

	go func() {
		var user models.UserModel
		if err := r.db.Where("id = ?", userId).Preload("Tasks").First(user).Error; err != nil {
			errorChan <- err
		}

		resultChan <- *user.ToUser()
	}()

	select {
	case user := <-resultChan:
		return user, nil
	case error := <-errorChan:
		return entities.User{}, error
	}
}

func (r *LocalUserRepository) SaveUser(user entities.User) *entities.User {
	resultChan := make(chan error)

	modelUser := models.UserModel{
		BaseModel: models.BaseModel{
			ID: uuid.NewV4(),
		},
		Name: user.Name,
	}

	go func() {
		if err := r.db.Create(&modelUser).Error; err != nil {
			resultChan <- err
		}

		resultChan <- nil
	}()

	<-resultChan

	return modelUser.ToUser()
}

func (r *LocalUserRepository) UpdateUser(user entities.User) error {
	resultChan := make(chan error)

	go func() {
		modelUser := models.UserModel{
			BaseModel: models.BaseModel{
				ID: uuid.NewV4(),
			},
			Name: user.Name,
		}

		if err := r.db.Save(&modelUser).Error; err != nil {
			resultChan <- err
		}

		resultChan <- nil
	}()

	return <-resultChan
}
