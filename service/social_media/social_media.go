package social_media

import (
	"errors"
	_entities "final-project-usamah/entities"
	_socmedRepository "final-project-usamah/repository/social_media"
	"time"
)

type SocmedService struct {
	socmedRepository _socmedRepository.SocmedRepositoryInterface
}

func NewSocmedService(socmedRepository _socmedRepository.SocmedRepositoryInterface) SocmedServiceInterface {
	return &SocmedService{
		socmedRepository: socmedRepository,
	}
}

func (ss *SocmedService) CreateSocmed(newSocmed _entities.Social_media, idToken int) (_entities.Social_media, error) {
	//validasi saat create sosmed
	if newSocmed.Name == "" {
		return newSocmed, errors.New("name is required")
	}
	if newSocmed.Social_media_url == "" {
		return newSocmed, errors.New("social_media_url is required")
	}

	newSocmed.User_id = idToken
	newSocmed.Created_at = time.Now()
	socmed, id, err := ss.socmedRepository.CreateSocmed(newSocmed)
	socmed.Id = id
	return socmed, err
}

func (ss *SocmedService) GetAllSocmed() ([]_entities.Social_media, error) {
	socmeds, err := ss.socmedRepository.GetAllSocmed()
	return socmeds, err
}

func (ss *SocmedService) UpdateSocmed(updateSocmed _entities.Social_media, idSocmed int, idToken int) (_entities.Social_media, error) {
	getSocmed, err := ss.socmedRepository.GetSocmed(idSocmed)
	if err != nil {
		return getSocmed, err
	}

	//validasi user login
	if idToken != getSocmed.User_id {
		return getSocmed, errors.New("unauthorized")
	}

	//validasi update social_media
	if updateSocmed.Name != "" {
		getSocmed.Name = updateSocmed.Name
	}
	if updateSocmed.Social_media_url != "" {
		getSocmed.Social_media_url = updateSocmed.Social_media_url
	}

	socmed, err := ss.socmedRepository.UpdateSocmed(getSocmed, idSocmed)
	socmed.Id = idSocmed
	socmed.Updated_at.Time = time.Now()
	return socmed, err
}

func (ss *SocmedService) DeleteSocmed(idSocmed int, idToken int) error {
	socmed, errGetSocmed := ss.socmedRepository.GetSocmed(idSocmed)
	if errGetSocmed != nil {
		return errGetSocmed
	}

	if idToken != socmed.User_id {
		return errors.New("unauthorized")
	}

	err := ss.socmedRepository.DeleteSocmed(idSocmed)
	return err
}
