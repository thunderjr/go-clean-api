package features

import (
	"fmt"
	"os"
	"testing"

	"github.com/thunderjr/go-clean-api/src/domain/repositories"
	infra_database "github.com/thunderjr/go-clean-api/src/infra/database"
	local_repositories "github.com/thunderjr/go-clean-api/src/infra/database/local/repositories"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() (*repositories.UserRepository, error) {
	db, err := gorm.Open(sqlite.Open("test_database.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = infra_database.Migrate(db)
	if err != nil {
		return nil, err
	}

	repo := local_repositories.NewLocalUserRepository(db)
	return &repo, nil
}

func makeUtils(r *repositories.UserRepository) *CreateUserFeature {
	return NewCreateUserFeature(r)
}

func makeSut(r *repositories.UserRepository) *GetUserFeature {
	return &GetUserFeature{
		repository: r,
	}
}

func TestGetUser(t *testing.T) {
	repository, err := initDB()
	if err != nil {
		t.Error("Unable to connect to Test Database", err)
	}
	defer func() {
		err = os.Remove("test_database.db")
		if err != nil {
			t.Error("Unable to delete test_database.db file", err)
		}
	}()

	createUserFeature := makeUtils(repository)
	sut := makeSut(repository)

	userName := "some-name"
	createdUser, err := createUserFeature.CreateUser(userName)
	if err != nil {
		t.Error("Error creating test user")
	}

	user, err := sut.GetUser(createdUser.ID)
	if err != nil {
		t.Error("Expected no error but got " + err.Error())
		return
	}

	fmt.Println(user)

	if user.ID != createdUser.ID {
		t.Errorf("Expected user with ID %s got ID: %s", createdUser.ID, user.ID)
		return
	}

	if user.Name != createdUser.Name {
		t.Errorf("Expected user with Name %s got Name: %s", createdUser.Name, user.Name)
		return
	}
}
