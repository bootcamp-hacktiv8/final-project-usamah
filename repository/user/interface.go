package user

import (
	"context"
	_entities "final-project-usamah/entities"
)

type UserRepositoryInterface interface {
	Register(ctx context.Context, user _entities.User) (_entities.User, int, error)
	Login(ctx context.Context, email string) (_entities.User, error)
	GetUser(ctx context.Context, idToken int) (_entities.User, error)
	UpdateUser(ctx context.Context, updateUser _entities.User, idToken int) (_entities.User, error)
	DeleteUser(ctx context.Context, idToken int) error
}
