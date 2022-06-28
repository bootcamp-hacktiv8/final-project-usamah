package social_media

import (
	"context"
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

func (ss *SocmedService) CreateSocmed(ctx context.Context, newSocmed _entities.Social_media, idToken int) (_entities.Social_media, error) {
	//validasi saat create sosmed
	if newSocmed.Name == "" {
		return newSocmed, errors.New("name is required")
	}
	if newSocmed.Social_media_url == "" {
		return newSocmed, errors.New("social_media_url is required")
	}

	newSocmed.User_id = idToken
	newSocmed.Created_at = time.Now()
	socmed, id, err := ss.socmedRepository.CreateSocmed(ctx, newSocmed)
	socmed.Id = id
	return socmed, err
}

func (ss *SocmedService) GetAllSocmed(ctx context.Context) ([]_entities.Social_media, error) {
	socmeds, err := ss.socmedRepository.GetAllSocmed(ctx)
	return socmeds, err
}

func (ss *SocmedService) UpdateSocmed(ctx context.Context, updateSocmed _entities.Social_media, idSocmed int, idToken int) (_entities.Social_media, error) {
	getSocmed, err := ss.socmedRepository.GetSocmed(ctx, idSocmed)
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

	socmed, err := ss.socmedRepository.UpdateSocmed(ctx, getSocmed, idSocmed)
	socmed.Id = idSocmed
	socmed.Updated_at.Time = time.Now()
	return socmed, err
}

func (ss *SocmedService) DeleteSocmed(ctx context.Context, idSocmed int, idToken int) error {
	socmed, errGetSocmed := ss.socmedRepository.GetSocmed(ctx, idSocmed)
	if errGetSocmed != nil {
		return errGetSocmed
	}

	if idToken != socmed.User_id {
		return errors.New("unauthorized")
	}

	err := ss.socmedRepository.DeleteSocmed(ctx, idSocmed)
	return err
}
