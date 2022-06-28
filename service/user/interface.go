package user

import (
	"context"
	"final-project-usamah/delivery/helper"
	_entities "final-project-usamah/entities"
)

type UserServiceInterface interface {
	Register(ctx context.Context, user _entities.User) (_entities.User, error)
	Login(ctx context.Context, inputLogin helper.LoginInput) (_entities.User, error)
	GetUser(ctx context.Context, idToken int) (_entities.User, error)
	UpdateUser(ctx context.Context, updateUser _entities.User, idToken int) (_entities.User, error)
	DeleteUser(ctx context.Context, idToken int) error
}
