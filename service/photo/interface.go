package photo

import (
	_entities "final-project-usamah/entities"
)

type PhotoServiceInterface interface {
	CreatePhoto(newPhoto _entities.Photo, idToken int) (_entities.Photo, error)
	GetAllPhoto() ([]_entities.Photo, error)
	UpdatePhoto(updatePhoto _entities.Photo, idPhoto int, idToken int) (_entities.Photo, error)
	DeletePhoto(idPhoto int, idToken int) error
}
