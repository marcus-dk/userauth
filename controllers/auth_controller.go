package controllers

import (
	"log"
	"userauth/database"
	"userauth/models"
	"userauth/utils"
	"golang.org/x/crypto/bcrypt"
)

/*func RegisterUser(user models.User) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	user.Password = string(hashedPassword)
	return database.AddUser(user)
}*/

func RegisterUser(user models.User) (models.User, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Error hashing password: %v", err)
        return models.User{}, err
    }
    user.Password = string(hashedPassword)
    log.Printf("Hashed password length: %d", len(user.Password))
    return database.AddUser(user)
}

func AuthenticateUser(email, password string) (string, error) {
    user, err := database.GetUserByEmail(email)
    if err != nil {
        log.Printf("Error retrieving user: %v", err)
        return "", err
    }
    log.Printf("Retrieved user: %+v", user)

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        log.Printf("Password comparison failed: %v", err)
        return "", err
    }

    return utils.GenerateJWT(user.Email)
}