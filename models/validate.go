package models

import (
	"errors"
)

var (
	ErrRequiredFirstName = errors.New("Name is required")
	ErrRequiredLastName  = errors.New("Last Name is required")
	ErrRequiredEmail     = errors.New("Email is required")
	ErrRequiredPassword  = errors.New("Password is required")
)

func IsEmpty(attr string) bool {
	if attr == "" {
		return true
	}
	return false
}

func ValidateNewUser(user User) (User, error) {
	if IsEmpty(user.FirstName) {
		return user, ErrRequiredFirstName
	}
	if IsEmpty(user.LastName) {
		return user, ErrRequiredLastName
	}
	if IsEmpty(user.Email) {
		return user, ErrRequiredEmail
	}
	if IsEmpty(user.Password) {
		return user, ErrRequiredPassword
	}
	return user, nil
}
