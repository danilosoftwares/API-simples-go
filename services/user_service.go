package services

import (
	"apisimples/models"
	"errors"
)

var nextID = 4

// Retorna todos os usuários
func GetAllUsers() []models.User {
	return models.Users
}

// Retorna um usuário por ID
func GetUserByID(id int) (models.User, error) {
	for _, user := range models.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, errors.New("usuário não encontrado")
}

// Cria um novo usuário
func CreateUser(user models.User) models.User {
	user.ID = nextID
	nextID++
	models.Users = append(models.Users, user)
	return user
}

// Atualiza um usuário existente
func UpdateUser(id int, updatedUser models.User) (models.User, error) {
	for i, user := range models.Users {
		if user.ID == id {
			updatedUser.ID = id
			models.Users[i] = updatedUser
			return updatedUser, nil
		}
	}
	return models.User{}, errors.New("usuário não encontrado")
}

// Remove um usuário por ID
func DeleteUser(id int) error {
	for i, user := range models.Users {
		if user.ID == id {
			models.Users = append(models.Users[:i], models.Users[i+1:]...)
			return nil
		}
	}
	return errors.New("usuário não encontrado")
}
