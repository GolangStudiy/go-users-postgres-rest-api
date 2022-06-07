package tests

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/GolangStudiy/go-users-postgres-rest-api/src/infrastructure"
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

type databaseClass struct {
	Datname string
}

func TestShouldBeReturnConnection(t *testing.T) {
	beforeTests()

	if connection == nil {
		t.Errorf("Connection cannot be null")
	}
}

func TestShouldBeReturnAllDatabaseNames(t *testing.T) {
	beforeTests()

	rows, err := infrastructure.RunQuery("SELECT datname FROM pg_database")

	if err != nil {
		t.Errorf("Query should to return data")
	}

	databases := make([]databaseClass, 0)
	var database databaseClass
	for rows.Next() {
		err := rows.Scan(&database.Datname)
		databases = append(databases, database)
		if err != nil {
			t.Errorf("Query should to return data")
		}
	}

	if len(databases) == 0 {
		t.Errorf("Query should to return data")
	}
}
