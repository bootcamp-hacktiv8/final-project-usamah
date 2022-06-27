package photo

import (
	response "final-project-usamah/delivery/helper/response/photo"
	_entities "final-project-usamah/entities"
)

type PhotoRepositoryInterface interface {
	CreatePhoto(newPhoto _entities.Photo) (_entities.Photo, int, error)
	GetAllPhoto() ([]response.GetPhotoFormatter, error)
	GetPhoto(idPhoto int) (_entities.Photo, error)
	UpdatePhoto(updatePhoto _entities.Photo, idPhoto int) (_entities.Photo, error)
	DeletePhoto(idPhoto int) error
}
