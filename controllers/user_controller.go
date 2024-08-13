package controllers

import (
	"errors"
	"userauth/database"
	"userauth/models"
)

var ErrUserNotFound = errors.New("user not found")

func GetAllUsers() ([]models.User, error) {
	return database.GetUsers()
}

func GetUserByID(id uint) (models.User, error) {
	return database.GetUserByID(id)
}

func CreateUser(user models.User) (models.User, error) {
	return database.AddUser(user)
}

func UpdateUser(id uint, updatedUser models.User) (models.User, error) {
    user, err := database.GetUserByID(id)
    if err != nil {
        if err == database.ErrUserNotFound {
            return models.User{}, ErrUserNotFound
        }
        return models.User{}, err
    }

    // Update fields if provided
    if updatedUser.Name != "" {
        user.Name = updatedUser.Name
    }
    if updatedUser.Email != "" {
        user.Email = updatedUser.Email
    }
    // Note: Password update should be handled separately for security reasons

    return database.UpdateUser(id, user)
}

func DeleteUser(id uint) error {
    err := database.DeleteUser(id)
    if err != nil {
        if err == database.ErrUserNotFound {
            return ErrUserNotFound
        }
        return err
    }
    return nil
}
