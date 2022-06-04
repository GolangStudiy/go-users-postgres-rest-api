package server

/// Go fmt import
import (
	"fmt"

	"github.com/GolangStudiy/go-users-postgres-rest-api/db"
)

// Go main function
func Main() {
	db.Migrate()
	msg := "App Started"
	fmt.Print(msg)
}
