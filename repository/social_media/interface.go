package social_media

import (
	response "final-project-usamah/delivery/helper/response/social_media"
	_entities "final-project-usamah/entities"
)

type SosmedRepositoryInterface interface {
	CreateSosmed(newSosmed _entities.Social_media) (_entities.Social_media, int, error)
	GetAllSosmed() ([]response.FormatGetSosmed, error)
	GetSosmed(idSosmed int) (_entities.Social_media, error)
	UpdateSosmed(updateSosmed _entities.Social_media, idSosmed int) (_entities.Social_media, error)
	DeleteSosmed(idSosmed int) error
}
