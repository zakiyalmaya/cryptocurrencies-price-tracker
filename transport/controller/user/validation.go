package user

import (
	"fmt"

	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

func validateUser(user *model.UserRequest) error {
	if user.Name == "" {
		return fmt.Errorf("name is required!")
	}

	if user.Username == "" {
		return fmt.Errorf("username is required!")
	}

	if user.Email == "" {
		return fmt.Errorf("email is required!")
	}

	if user.Password == "" {
		return fmt.Errorf("password is required!")
	}

	if user.ConfirmPassword != user.Password {
		return fmt.Errorf("password and confirm password does not match!")
	}

	return nil
}