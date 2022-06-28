package photo

import (
	"context"
	_entities "final-project-usamah/entities"
)

type PhotoRepositoryInterface interface {
	CreatePhoto(ctx context.Context, newPhoto _entities.Photo) (_entities.Photo, int, error)
	GetAllPhoto(ctx context.Context) ([]_entities.Photo, error)
	GetPhoto(ctx context.Context, idPhoto int) (_entities.Photo, error)
	UpdatePhoto(ctx context.Context, updatePhoto _entities.Photo, idPhoto int) (_entities.Photo, error)
	DeletePhoto(ctx context.Context, idPhoto int) error
}
