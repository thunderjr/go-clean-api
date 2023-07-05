package repositories

import (
	"github.com/thunderjr/go-clean-api/src/domain/entities"
)

type UserRepository interface {
	GetUser(userId string) (entities.User, error)
	SaveUser(user entities.User) error
	UpdateUser(user entities.User) error
}
