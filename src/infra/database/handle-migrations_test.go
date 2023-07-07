package infra_database_test

import (
	"os"
	"testing"

	infra_database "github.com/thunderjr/go-clean-api/src/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMigrations(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test_database.db"), &gorm.Config{})
	if err != nil {
		t.Error("Unable to connect to Test Database", err)
	}

	defer func() {
		err = os.Remove("test_database.db")
		if err != nil {
			t.Error("Unable to delete test_database.db file", err)
		}
	}()

	err = infra_database.Migrate(db)
	if err != nil {
		t.Error("Error migrating database: ", err)
	}
}
