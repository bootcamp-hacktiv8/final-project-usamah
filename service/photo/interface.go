package photo

import (
	"context"
	_entities "final-project-usamah/entities"
)

type PhotoServiceInterface interface {
	CreatePhoto(ctx context.Context, newPhoto _entities.Photo, idToken int) (_entities.Photo, error)
	GetAllPhoto(ctx context.Context) ([]_entities.Photo, error)
	UpdatePhoto(ctx context.Context, updatePhoto _entities.Photo, idPhoto int, idToken int) (_entities.Photo, error)
	DeletePhoto(ctx context.Context, idPhoto int, idToken int) error
}
