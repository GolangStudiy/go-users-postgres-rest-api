package infrastructure

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID `gorm:"primaryKey;column:id"`
	Email string    `gorm:"column:email"`
}
