package services

import (
	domain "github.com/GolangStudiy/go-users-postgres-rest-api/src/domain/user"
	infrastructure "github.com/GolangStudiy/go-users-postgres-rest-api/src/infrastructure/user"
)

func Post(email string) (string, error) {
	email, err := domain.ValidateEmail(email)

	if err != nil {
		return "", err
	}

	email, err = domain.UserRepository.Post(
		&infrastructure.UserRepository{},
		domain.User{Email: email},
	)

	return email, err
}

func GetIdByEmail(email string) (string, error) {
	id, err := domain.UserRepository.GetIdByEmail(
		&infrastructure.UserRepository{},
		email,
	)

	return id, err
}
