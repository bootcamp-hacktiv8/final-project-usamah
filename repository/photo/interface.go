package photo

import (
	_entities "final-project-usamah/entities"
)

type PhotoRepositoryInterface interface {
	CreatePhoto(newPhoto _entities.Photo) (_entities.Photo, int, error)
	GetAllPhoto() ([]_entities.Photo, error)
	GetPhoto(idPhoto int) (_entities.Photo, error)
	UpdatePhoto(updatePhoto _entities.Photo, idPhoto int) (_entities.Photo, error)
	DeletePhoto(idPhoto int) error
}
