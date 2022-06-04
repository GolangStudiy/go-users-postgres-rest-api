package tests

import (
	"database/sql"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	databaseclient "github.com/GolangStudiy/go-users-postgres-rest-api/src/configurations"
	"github.com/GolangStudiy/go-users-postgres-rest-api/src/server"
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

	connection = databaseclient.GetConnection()
}

func TestShouldBePrintCorrectMessage(t *testing.T) {
	beforeTests(t)
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	server.Main()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if string(out) != "App Started" {
		t.Errorf("Expected %s, got %s", "App Started", out)
	}
}
