package user

import (
	"github.com/google/uuid"
)

type UserRepository interface {
	Post(user User) string
	GetIdByEmail(email string) uuid.UUID
}
