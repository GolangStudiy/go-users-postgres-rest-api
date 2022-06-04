package databaseclient

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "users"
)

func GetConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	connection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Fail to get database connection : %s", err)
	}

	err = connection.Ping()
	if err != nil {
		log.Fatalf("Error DB Ping : %s", err)
	}

	return connection
}

func RunQuery(sql string) *sql.Rows {
	connection := GetConnection()
	data, err := connection.Query(sql)
	defer connection.Close()
	if err != nil {
		log.Printf("Error when execute query : %s", err)
		return nil
	}

	return data
}
