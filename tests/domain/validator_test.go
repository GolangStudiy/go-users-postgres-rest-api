package tests

import (
	"testing"

	domain "github.com/GolangStudiy/go-users-postgres-rest-api/src/domain/user"
)

func TestShouldBeReturnErrorWhenEmailIsWrong(t *testing.T) {
	address, err := domain.ValidateEmail("foobar")
	if err == nil || address != "" {
		t.Errorf("Should be return an error when the e-mail is in incorrect format")
	}

	address, err = domain.ValidateEmail("foobar@")
	if err == nil || address != "" {
		t.Errorf("Should be return an error when the e-mail is in incorrect format")
	}

	address, err = domain.ValidateEmail("foobar@.com")
	if err == nil || address != "" {
		t.Errorf("Should be return an error when the e-mail is in incorrect format")
	}

	address, err = domain.ValidateEmail("@.com")
	if err == nil || address != "" {
		t.Errorf("Should be return an error when the e-mail is in incorrect format")
	}
}

func TestShouldBeReturnReturnAddressIfEmailIsCorrect(t *testing.T) {
	address, err := domain.ValidateEmail("john-doe@email.com")
	if err != nil || address == "" {
		t.Errorf("Should not return an error when the e-mail is in correct format")
	}
}
