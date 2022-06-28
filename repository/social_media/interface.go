package social_media

import (
	_entities "final-project-usamah/entities"
)

type SocmedRepositoryInterface interface {
	CreateSocmed(newSocmed _entities.Social_media) (_entities.Social_media, int, error)
	GetAllSocmed() ([]_entities.Social_media, error)
	GetSocmed(idSocmed int) (_entities.Social_media, error)
	UpdateSocmed(updateSocmed _entities.Social_media, idSocmed int) (_entities.Social_media, error)
	DeleteSocmed(idSocmed int) error
}
