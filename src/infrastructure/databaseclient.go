package infrastructure

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func GetDbConnection() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	connection, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
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
		return nil, err
	}

	return connection, nil
}

func RunQuery(sql string) (*sql.Rows, error) {
	connection, err := GetDbConnection()

	if err != nil {
		return nil, err
	}

	data, err := connection.Query(sql)
	defer connection.Close()

	if err != nil {
		return nil, err
	}

	return data, nil
}
