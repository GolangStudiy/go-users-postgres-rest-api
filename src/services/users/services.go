package services

import (
	domain "github.com/GolangStudiy/go-users-postgres-rest-api/src/domain/user"
	entrypoint "github.com/GolangStudiy/go-users-postgres-rest-api/src/entrypoint/user"
	infrastructure "github.com/GolangStudiy/go-users-postgres-rest-api/src/infrastructure/user"
)

func Post(requestUser entrypoint.RequestUser) (entrypoint.ResponseUser, error) {
	email, err := domain.UserRepository.Post(
		&infrastructure.UserRepository{},
		domain.User{Email: requestUser.Email},
	)

	return entrypoint.ResponseUser{Email: email}, err
}

func GetIdByEmail(email string) (string, error) {
	id, err := domain.UserRepository.GetIdByEmail(
		&infrastructure.UserRepository{},
		email,
	)

	return id, err
}
