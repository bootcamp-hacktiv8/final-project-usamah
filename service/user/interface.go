package user

import (
	"final-project-usamah/delivery/helper"
	_entities "final-project-usamah/entities"
)

type UserServiceInterface interface {
	Register(user _entities.User) (_entities.User, error)
	Login(inputLogin helper.LoginInput) (_entities.User, error)
	GetUser(idToken int) (_entities.User, error)
	UpdateUser(updateUser _entities.User, idToken int) (_entities.User, error)
	DeleteUser(idToken int) error
}
