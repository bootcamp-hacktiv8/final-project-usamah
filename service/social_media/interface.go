package social_media

import (
	_entities "final-project-usamah/entities"
)

type SocmedServiceInterface interface {
	CreateSocmed(newSocmed _entities.Social_media, idToken int) (_entities.Social_media, error)
	GetAllSocmed() ([]_entities.Social_media, error)
	UpdateSocmed(updateSocmed _entities.Social_media, idSocmed int, idToken int) (_entities.Social_media, error)
	DeleteSocmed(idSocmed int, idToken int) error
}
