package main

/// Go fmt import
import (
	"fmt"

	databaseclient "github.com/GolangStudiy/go-users-postgres-rest-api/databaseclient"
)

// Go main function
func main() {
	msg := "App Started"
	databaseclient.GetConnection()
	fmt.Print(msg)
}
