package social_media

import (
	"errors"
	response "final-project-usamah/delivery/helper/response/social_media"
	_entities "final-project-usamah/entities"
	_sosmedRepository "final-project-usamah/repository/social_media"
	"time"
)

type SosmedService struct {
	sosmedRepository _sosmedRepository.SosmedRepositoryInterface
}

func NewSosmedService(sosmedRepository _sosmedRepository.SosmedRepositoryInterface) SosmedServiceInterface {
	return &SosmedService{
		sosmedRepository: sosmedRepository,
	}
}

func (ss *SosmedService) CreateSosmed(newSosmed _entities.Social_media, idToken int) (_entities.Social_media, error) {
	//validasi saat create sosmed
	if newSosmed.Name == "" {
		return newSosmed, errors.New("name is required")
	}
	if newSosmed.Social_media_url == "" {
		return newSosmed, errors.New("social_media_url is required")
	}

	newSosmed.User_id = idToken
	newSosmed.Created_at = time.Now()
	sosmed, id, err := ss.sosmedRepository.CreateSosmed(newSosmed)
	sosmed.Id = id
	return sosmed, err
}

func (ss *SosmedService) GetAllSosmed() ([]response.FormatGetSosmed, error) {
	sosmeds, err := ss.sosmedRepository.GetAllSosmed()
	return sosmeds, err
}

func (ss *SosmedService) UpdateSosmed(updateSosmed _entities.Social_media, idSosmed int, idToken int) (_entities.Social_media, error) {
	getSosmed, err := ss.sosmedRepository.GetSosmed(idSosmed)
	if err != nil {
		return getSosmed, err
	}

	//validasi user login
	if idToken != getSosmed.User_id {
		return getSosmed, errors.New("unauthorized")
	}

	//validasi update social_media
	if updateSosmed.Name != "" {
		getSosmed.Name = updateSosmed.Name
	}
	if updateSosmed.Social_media_url != "" {
		getSosmed.Social_media_url = updateSosmed.Social_media_url
	}

	sosmed, err := ss.sosmedRepository.UpdateSosmed(getSosmed, idSosmed)
	sosmed.Id = idSosmed
	sosmed.Updated_at.Time = time.Now()
	return sosmed, err
}

func (ss *SosmedService) DeleteSosmed(idSosmed int, idToken int) error {
	sosmed, errGetSosmed := ss.sosmedRepository.GetSosmed(idSosmed)
	if errGetSosmed != nil {
		return errGetSosmed
	}

	if idToken != sosmed.User_id {
		return errors.New("unauthorized")
	}

	err := ss.sosmedRepository.DeleteSosmed(idSosmed)
	return err
}
