package models

import (
	"errors"
)

var (
	ErrRequiredFirstName = errors.New("name is required")
	ErrRequiredLastName  = errors.New("nast Name is required")
	ErrRequiredEmail     = errors.New("email is required")
	ErrRequiredPassword  = errors.New("password is required")
)

func IsEmpty(attr string) bool {
	if attr == "" {
		return true
	}

	return false
}

func ValidateNewUser(user User) (User, error) {
	if IsEmpty(user.FirstName) {
		return User{}, ErrRequiredFirstName
	}
	if IsEmpty(user.LastName) {
		return User{}, ErrRequiredLastName
	}
	if IsEmpty(user.Email) {
		return User{}, ErrRequiredEmail
	}
	if IsEmpty(user.Password) {
		return User{}, ErrRequiredPassword
	}

	return user, nil
}
