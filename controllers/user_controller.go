package controllers

import (
	"userauth/database"
	"userauth/models"
)

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
	return database.UpdateUser(id, updatedUser)
}

func DeleteUser(id uint) error {
	return database.DeleteUser(id)
}
