package services

import (
	"balancer/contracts"
	"balancer/models"
	"balancer/repositories"
	utils "balancer/utils/encryption"
)

func LoginValidation(loginData *contracts.LoginRequest) error {
	user := models.User{
		Email: loginData.Email,
		Hash:  utils.GenerateHash(loginData.Password),
	}

	err := repositories.LoginValidation(user)
	if err != nil {
		return err
	}

	return nil
}
