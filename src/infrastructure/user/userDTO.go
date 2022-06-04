package user

import "github.com/google/uuid"

type UserDTO struct {
	ID    uuid.UUID `gorm:"primaryKey;column:id"`
	Email string    `gorm:"column:email"`
}
