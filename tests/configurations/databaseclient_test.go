package tests

import (
	"database/sql"
	"os"
	"strconv"
	"testing"

	"github.com/GolangStudiy/go-users-postgres-rest-api/src/configurations"
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

type databaseClass struct {
	Datname string
}

func TestShouldBeReturnConnection(t *testing.T) {
	beforeTests(t)

	if connection == nil {
		t.Errorf("Connection cannot be null")
	}
}

func TestShouldBeReturnAllDatabaseNames(t *testing.T) {
	beforeTests(t)

	rows := configurations.RunQuery("SELECT datname FROM pg_database")

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
