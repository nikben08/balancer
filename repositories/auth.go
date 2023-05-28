package repositories

import (
	"balancer/models"
	"errors"
)

func LoginValidation(user models.User) error {
	users, err := DB.Query("SELECT * FROM `users` WHERE email='" + user.Email + "' AND hash='" + user.Hash + "'")
	if err != nil {
		return err
	}

	count := 0
	for users.Next() {
		count++
	}

	if count == 0 {
		return errors.New("Wrong E-mail or password")
	}

	return nil
}
