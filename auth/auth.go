package auth

import (
	"errors"

	"github.com/ssr0016/webapp/models"
	"github.com/ssr0016/webapp/utils"
)

var (
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")
)

func Signin(email, password string) (models.User, error) {
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, ErrInvalidEmail
	}
	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		return models.User{}, ErrInvalidPassword
	}
	return user, nil
}
