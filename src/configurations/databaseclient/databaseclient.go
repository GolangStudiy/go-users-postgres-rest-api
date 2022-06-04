package databaseclient

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func getConnection() *sql.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	connection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Fail to get database connection : %s", err)
	}

	err = connection.Ping()
	limit := 0
	for limit < 10 && err != nil {
		time.Sleep(2 * time.Second)
		err = connection.Ping()
		if err == nil {
			limit = 10
		}
	}
	if err != nil {
		log.Fatalf("Error DB Ping : %s", err)
	}

	return connection
}

func RunQuery(sql string) *sql.Rows {
	connection := getConnection()
	data, err := connection.Query(sql)
	defer connection.Close()
	if err != nil {
		log.Printf("Error when execute query : %s", err)
		return nil
	}

	return data
}
