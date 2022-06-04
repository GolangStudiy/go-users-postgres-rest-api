package user

import (
	"fmt"

	databaseclient "github.com/GolangStudiy/go-users-postgres-rest-api/src/configurations"
	"github.com/GolangStudiy/go-users-postgres-rest-api/src/domain/user"
	"github.com/google/uuid"
)

type UserRepository struct {
}

func (r *UserRepository) Post(user user.User) string {
	userDTO := UserDTO{
		Email: user.Email,
	}
	query := fmt.Sprintf(`INSERT INTO users (email) VALUES ('%s') RETURNING email ; `, userDTO.Email)
	rows := databaseclient.RunQuery(query)
	var email string
	for rows.Next() {
		rows.Scan(&email)
	}

	return email
}

func (r *UserRepository) GetIdByEmail(email string) uuid.UUID {
	query := fmt.Sprintf(`SELECT id FROM users WHERE email = '%s' ;`, email)
	rows := databaseclient.RunQuery(query)
	var id uuid.UUID
	for rows.Next() {
		rows.Scan(&id)
	}

	return id
}
