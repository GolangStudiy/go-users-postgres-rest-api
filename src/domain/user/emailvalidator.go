package domain

import "net/mail"

func ValidateEmail(email string) (string, error) {
	address, err := mail.ParseAddress(email)

	if err != nil {
		return "", err
	}

	return address.Address, err
}
