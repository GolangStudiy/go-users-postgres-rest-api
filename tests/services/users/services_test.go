package tests

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/GolangStudiy/go-users-postgres-rest-api/db"
	"github.com/GolangStudiy/go-users-postgres-rest-api/src/infrastructure"
	services "github.com/GolangStudiy/go-users-postgres-rest-api/src/services/user"
	"github.com/GolangStudiy/go-users-postgres-rest-api/tests"
)

var connection *sql.DB

func beforeTests() {
	db := tests.MountDatabaseContainer()
	dbPort := db.GetPort()

	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", strconv.Itoa(dbPort))
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_NAME", "users")

	var err error
	connection, err = infrastructure.GetDbConnection()

	if err != nil {
		log.Fatal(err)
	}
}

func TestShouldBeCreateUserAndReturnTheEmail(t *testing.T) {
	beforeTests()
	db.Migrate()

	email := "john@doe.com"
	responseEmail, err := services.Post(email)

	if err != nil {
		t.Errorf("Should not be have error here")
	}

	if responseEmail != email {
		t.Errorf("Expected %s, got %s", email, responseEmail)
	}
}

func TestShouldBeCreateUserAndReturnTheUUID(t *testing.T) {
	beforeTests()
	db.Migrate()

	email := "john@doe.com"
	responseEmail, err := services.Post(email)

	if err != nil {
		t.Errorf("Should not be have error here")
	}

	id, err := services.GetIdByEmail(responseEmail)

	if err != nil {
		t.Errorf("Should not be have error here")
	}

	if id == "" {
		t.Errorf("Expected something different of null")
	}
}

func TestShouldBeReturnNullIfTryToPostTheSameEmail(t *testing.T) {
	beforeTests()
	db.Migrate()

	email := "john@doe.com"
	responseEmail, err := services.Post(email)

	if err != nil {
		t.Errorf("Should not be have error here")
	}

	responseEmail, err = services.Post(string(responseEmail))

	if err == nil || !strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		t.Errorf(err.Error())
	}
}
