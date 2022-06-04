package migrate

import (
	"log"
	"path"
	"runtime"

	databaseclient "github.com/GolangStudiy/go-users-postgres-rest-api/src/configurations"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Main() {
	connection := databaseclient.GetConnection()
	driver, err := postgres.WithInstance(connection, &postgres.Config{})

	if err != nil {
		log.Fatal(err)
	}

	_, b, _, _ := runtime.Caller(0)
	projectDir := path.Join(path.Dir(b))

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+projectDir+"/migrations",
		"users", driver)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
