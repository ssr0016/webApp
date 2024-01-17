package models

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

var (
	ErrRequiredFirstName = errors.New("First name is required")
	ErrRequiredLastName  = errors.New("Last name is required")
	ErrRequiredEmail     = errors.New("Email is required")
	ErrInvalidEmail      = errors.New("Invalid email")
	ErrRequiredPassword  = errors.New("Password is required")
	ErrMaxlimit          = errors.New("Exceeded the maximum character limit")
	ErrDuplicateKeyEmail = errors.New("Email already exists")
)

func IsEmpty(attr string) bool {
	if attr == "" {
		return true
	}
	return false
}

func Trim(attr string) string {
	return strings.TrimSpace(attr)
}

func IsEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return false
	}

	return true
}

func Max(attr string, lim int) bool {
	if len(attr) <= lim {
		return true
	}
	return false
}

func ValidateLimitFields(user User) (User, error) {
	if !Max(user.FirstName, 15) || !Max(user.LastName, 20) || !Max(user.Email, 40) || !Max(user.Password, 100) {
		return user, ErrMaxlimit
	}

	return user, nil
}

func UniqueEmail(email string) (bool, error) {
	con := Connect()
	defer con.Close()
	sql := "select count(email) from users where email = $1"
	rs, err := con.Query(sql, email)
	if err != nil {
		return false, err
	}
	defer rs.Close()
	var count int64
	if rs.Next() {
		err := rs.Scan(&count)
		if err != nil {
			return false, err
		}
	}
	if count > 0 {
		return false, ErrDuplicateKeyEmail
	}
	return true, nil
}

func ValidateNewUser(user User) (User, error) {
	_, err := UniqueEmail(user.Email)
	if err != nil {
		return User{}, err
	}

	user, err = ValidateLimitFields(user)
	if err != nil {
		return user, err
	}

	user.FirstName = Trim(user.FirstName)
	user.LastName = Trim(user.LastName)
	user.Email = Trim(strings.ToLower(user.Email))

	if IsEmpty(user.FirstName) {
		return User{}, ErrRequiredFirstName
	}
	if IsEmpty(user.LastName) {
		return User{}, ErrRequiredLastName
	}
	if IsEmpty(user.Email) {
		return User{}, ErrRequiredEmail
	}
	if !IsEmail(user.Email) {
		return User{}, ErrInvalidEmail
	}
	if IsEmpty(user.Password) {
		return User{}, ErrRequiredPassword
	}

	return user, nil
}
