package database

import (
	"encoding/json"
	"log"
	"errors"
	"io/ioutil"
	"os"
	"userauth/models"
	"sync"
)

var (
	filePath = "data/users.json"
	mutex    = &sync.Mutex{}
)

var ErrUserNotFound = errors.New("user not found")

func init() {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.MkdirAll("data", os.ModePerm)
		if err != nil {
			panic(err)
		}

		emptyData := []models.User{}
		data, _ := json.Marshal(emptyData)
		ioutil.WriteFile(filePath, data, 0644)
	}
}

func readFromFile() ([]models.User, error) {
    mutex.Lock()
    defer mutex.Unlock()

    file, err := ioutil.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    log.Printf("Read from file: %s", string(file))

    var users []models.User
    err = json.Unmarshal(file, &users)
    if err != nil {
        return nil, err
    }

    return users, nil
}

func writeToFile(users []models.User) error {
    mutex.Lock()
    defer mutex.Unlock()

    data, err := json.MarshalIndent(users, "", "  ")
    if err != nil {
        return err
    }

    log.Printf("Writing to file: %s", string(data))
    return ioutil.WriteFile(filePath, data, 0644)
}

func GetUsers() ([]models.User, error) {
	return readFromFile()
}

func AddUser(user models.User) (models.User, error) {
    log.Printf("Adding user: %+v", user)
    users, err := readFromFile()
    if err != nil {
        log.Printf("Error reading from file: %v", err)
        return models.User{}, err
    }

    for _, u := range users {
        if u.Email == user.Email {
            return models.User{}, errors.New("user already exists")
        }
    }

    if len(users) > 0 {
        user.ID = users[len(users)-1].ID + 1
    } else {
        user.ID = 1
    }

    users = append(users, user)
    err = writeToFile(users)
    if err != nil {
        log.Printf("Error writing to file: %v", err)
        return models.User{}, err
    }

    log.Printf("User added successfully: %+v", user)
    return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
    users, err := readFromFile()
    if err != nil {
        return models.User{}, err
    }

    for _, u := range users {
        if u.Email == email {
            log.Printf("Retrieved user with email: %s", email)
            log.Printf("Retrieved user with password hash length: %d", len(u.Password))
            return u, nil
        }
    }

    return models.User{}, ErrUserNotFound
}

func GetUserByID(id uint) (models.User, error) {
	users, err := readFromFile()
	if err != nil {
		return models.User{}, err
	}

	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}

	return models.User{}, errors.New("user not found")
}

func UpdateUser(id uint, updatedUser models.User) (models.User, error) {
	users, err := readFromFile()
	if err != nil {
		return models.User{}, err
	}

	for i, u := range users {
		if u.ID == id {
			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email
			users[i].Password = updatedUser.Password
			err = writeToFile(users)
			if err != nil {
				return models.User{}, err
			}
			return users[i], nil
		}
	}

	return models.User{}, errors.New("user not found")
}

func DeleteUser(id uint) error {
	users, err := readFromFile()
	if err != nil {
		return err
	}

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return writeToFile(users)
		}
	}

	return errors.New("user not found")
}
