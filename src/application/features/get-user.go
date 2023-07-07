package features

import (
	"github.com/thunderjr/go-clean-api/src/domain/entities"
	"github.com/thunderjr/go-clean-api/src/domain/repositories"
)

type GetUserFeature struct {
	repository *repositories.UserRepository
}

func NewGetUserFeature(r *repositories.UserRepository) *GetUserFeature {
	return &GetUserFeature{
		repository: r,
	}
}

func (f *GetUserFeature) GetUser(userId string) (*entities.User, error) {
	user, err := (*f.repository).GetUser(userId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
