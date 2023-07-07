package features

import (
	"github.com/thunderjr/go-clean-api/src/domain/entities"
	"github.com/thunderjr/go-clean-api/src/domain/repositories"
)

type CreateUserFeature struct {
	repository *repositories.UserRepository
}

func NewCreateUserFeature(r *repositories.UserRepository) *CreateUserFeature {
	return &CreateUserFeature{
		repository: r,
	}
}

func (f *CreateUserFeature) CreateUser(name string) (*entities.User, error) {
	user := entities.User{
		Name: name,
	}

	user = *(*f.repository).SaveUser(user)

	return &user, nil
}
