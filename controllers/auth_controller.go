package controllers

import (
	"userauth/database"
	"userauth/models"
	"userauth/utils"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user models.User) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	user.Password = string(hashedPassword)
	return database.AddUser(user)
}

func AuthenticateUser(email, password string) (string, error) {
	user, err := database.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	return utils.GenerateJWT(user.Email)
}
