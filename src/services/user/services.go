package services

import (
	domain "github.com/GolangStudiy/go-users-postgres-rest-api/src/domain/user"
	infrastructure "github.com/GolangStudiy/go-users-postgres-rest-api/src/infrastructure/user"
)

func Post(domainUser domain.User) (domain.User, error) {
	email, err := domain.UserRepository.Post(
		&infrastructure.UserRepository{},
		domainUser,
	)

	return domain.User{Email: email}, err
}

func GetIdByEmail(email string) (string, error) {
	id, err := domain.UserRepository.GetIdByEmail(
		&infrastructure.UserRepository{},
		email,
	)

	return id, err
}
