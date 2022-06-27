package social_media

import (
	_entities "final-project-usamah/entities"
)

type SosmedRepositoryInterface interface {
	CreateSosmed(newSosmed _entities.Social_media) (_entities.Social_media, int, error)
	GetAllSosmed() ([]_entities.Social_media, error)
	GetSosmed(idSosmed int) (_entities.Social_media, error)
	UpdateSosmed(updateSosmed _entities.Social_media, idSosmed int) (_entities.Social_media, error)
	DeleteSosmed(idSosmed int) error
}
