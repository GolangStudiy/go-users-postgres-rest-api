package main

/// Go fmt import
import (
	"fmt"

	migrate "github.com/GolangStudiy/go-users-postgres-rest-api/migrate"
)

// Go main function
func main() {
	migrate.Main()
	msg := "App Started"
	fmt.Print(msg)
}
