package social_media

import (
	_entities "final-project-usamah/entities"
)

type SosmedServiceInterface interface {
	CreateSosmed(newSosmed _entities.Social_media, idToken int) (_entities.Social_media, error)
	GetAllSosmed() ([]_entities.Social_media, error)
	UpdateSosmed(updateSosmed _entities.Social_media, idSosmed int, idToken int) (_entities.Social_media, error)
	DeleteSosmed(idSosmed int, idToken int) error
}
