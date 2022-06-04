package infrastructure

import (
	"fmt"

	"github.com/GolangStudiy/go-users-postgres-rest-api/src/configurations"
	domain "github.com/GolangStudiy/go-users-postgres-rest-api/src/domain/user"
)

type UserRepository struct {
}

func (r *UserRepository) Post(domainUser domain.User) (string, error) {
	rows, err := configurations.RunQuery(
		fmt.Sprintf(
			"INSERT INTO users (email) VALUES ('%s') RETURNING email;",
			User{Email: domainUser.Email}.Email,
		),
	)

	if err != nil {
		return "", err
	}

	var email string
	if rows.Next() {
		rows.Scan(&email)
	}

	return email, nil
}

func (r *UserRepository) GetIdByEmail(email string) (string, error) {
	rows, err := configurations.RunQuery(fmt.Sprintf("SELECT id FROM users WHERE email = '%s';", email))

	if err != nil {
		return "", nil
	}

	var id string
	if rows.Next() {
		rows.Scan(&id)
	}

	return id, nil
}
