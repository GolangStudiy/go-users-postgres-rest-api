package migrate

import (
	"database/sql"
	"os"
	"strconv"
	"testing"

	databaseclient "github.com/GolangStudiy/go-users-postgres-rest-api/databaseclient"
	postgrescontainer "github.com/GolangStudiy/go-users-postgres-rest-api/postgrescontainer"
)

var connection *sql.DB

func beforeTests(t *testing.T) {
	db := postgrescontainer.MountDatabaseContainer(t)
	dbPort := db.GetPort(t)

	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", strconv.Itoa(dbPort))
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_NAME", "users")

	connection = databaseclient.GetConnection()
}

type tableClass struct {
	table_name string
}

func TestShouldBeCreateUsersTable(t *testing.T) {
	beforeTests(t)
	Main()

	rows := databaseclient.RunQuery("SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE' AND table_name='users'")

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
