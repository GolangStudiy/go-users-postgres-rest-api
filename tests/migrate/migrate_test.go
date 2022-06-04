package tests

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/GolangStudiy/go-users-postgres-rest-api/db"
	"github.com/GolangStudiy/go-users-postgres-rest-api/src/configurations"
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
	connection, err = configurations.GetDbConnection()

	if err != nil {
		log.Fatal(err)
	}
}

type tableClass struct {
	table_name string
}

func TestShouldBeCreateUsersTable(t *testing.T) {
	beforeTests()
	db.Migrate()

	rows, err := configurations.RunQuery("SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE' AND table_name='users'")

	if err != nil {
		t.Errorf("Query should to return data")
	}

	var table tableClass
	for rows.Next() {
		err := rows.Scan(&table.table_name)
		if err != nil {
			t.Errorf("The users table should be exists")
		}
	}

	if table.table_name != "users" {
		t.Errorf("The users table should be exists")
	}
}
