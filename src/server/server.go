package server

/// Go fmt import
import (
	"fmt"

	"github.com/GolangStudiy/go-users-postgres-rest-api/migrate"
)

// Go main function
func Main() {
	migrate.Main()
	msg := "App Started"
	fmt.Print(msg)
}
