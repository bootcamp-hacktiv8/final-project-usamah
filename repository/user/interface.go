package user

import (
	_entities "final-project-usamah/entities"
)

type UserRepositoryInterface interface {
	Register(user _entities.User) (_entities.User, int, error)
	Login(email string) (_entities.User, error)
	GetUser(idToken int) (_entities.User, error)
	UpdateUser(updateUser _entities.User, idToken int) (_entities.User, error)
	DeleteUser(idToken int) error
}
