package user

import (
	_request "final-project-usamah/delivery/helper/request/user"
	_entities "final-project-usamah/entities"
)

type UserServiceInterface interface {
	Register(user _entities.User) (_entities.User, error)
	Login(inputLogin _request.FormatLogin) (_entities.User, error)
	GetUser(idToken int) (_entities.User, error)
	UpdateUser(updateUser _entities.User, idToken int) (_entities.User, error)
	DeleteUser(idToken int) error
}
