package tests

import (
	"database/sql"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/GolangStudiy/go-users-postgres-rest-api/db"
	"github.com/GolangStudiy/go-users-postgres-rest-api/src/configurations"
	entrypoint "github.com/GolangStudiy/go-users-postgres-rest-api/src/entrypoint/user"
	services "github.com/GolangStudiy/go-users-postgres-rest-api/src/services/users"
	"github.com/GolangStudiy/go-users-postgres-rest-api/tests"
)

var connection *sql.DB

func beforeTests(t *testing.T) {
	db := tests.MountDatabaseContainer(t)
	dbPort := db.GetPort(t)

	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", strconv.Itoa(dbPort))
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_NAME", "users")

	connection = configurations.GetDbConnection()
}

func TestShouldBeCreateUserAndReturnTheEmail(t *testing.T) {
	beforeTests(t)
	db.Migrate()

	email := "john@doe.com"
	responseUser, err := services.Post(entrypoint.RequestUser{Email: email})

	if err != nil {
		t.Errorf("Should not be have error here")
	}

	if responseUser.Email != email {
		t.Errorf("Expected %s, got %s", email, responseUser.Email)
	}
}

func TestShouldBeCreateUserAndReturnTheUUID(t *testing.T) {
	beforeTests(t)
	db.Migrate()

	email := "john@doe.com"
	responseUser, err := services.Post(entrypoint.RequestUser{Email: email})

	if err != nil {
		t.Errorf("Should not be have error here")
	}

	id, err := services.GetIdByEmail(responseUser.Email)

	if err != nil {
		t.Errorf("Should not be have error here")
	}

	if id == "" {
		t.Errorf("Expected something different of null")
	}
}

func TestShouldBeReturnNullIfTryToPostTheSameEmail(t *testing.T) {
	beforeTests(t)
	db.Migrate()

	email := "john@doe.com"
	responseUser, err := services.Post(entrypoint.RequestUser{Email: email})

	if err != nil {
		t.Errorf("Should not be have error here")
	}

	responseUser, err = services.Post(entrypoint.RequestUser{Email: string(responseUser.Email)})

	if err == nil || !strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		t.Errorf(err.Error())
	}
}
