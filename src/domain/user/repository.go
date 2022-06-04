package domain

type UserRepository interface {
	Post(user User) (string, error)
	GetIdByEmail(email string) (string, error)
}
