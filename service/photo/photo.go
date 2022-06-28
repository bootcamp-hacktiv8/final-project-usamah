package photo

import (
	"context"
	"errors"
	_entities "final-project-usamah/entities"
	_photoRepository "final-project-usamah/repository/photo"
	"time"
)

type PhotoService struct {
	photoRepository _photoRepository.PhotoRepositoryInterface
}

func NewPhotoService(photoRepository _photoRepository.PhotoRepositoryInterface) PhotoServiceInterface {
	return &PhotoService{
		photoRepository: photoRepository,
	}
}

func (ps *PhotoService) CreatePhoto(ctx context.Context, newPhoto _entities.Photo, idToken int) (_entities.Photo, error) {
	//validasi saat create photo
	if newPhoto.Title == "" {
		return newPhoto, errors.New("title is required")
	}
	if newPhoto.Photo_url == "" {
		return newPhoto, errors.New("photo_url is required")
	}

	newPhoto.User_id = idToken
	photo, id, err := ps.photoRepository.CreatePhoto(ctx, newPhoto)
	photo.Id = id
	photo.Created_at = time.Now()
	return photo, err
}

func (ps *PhotoService) GetAllPhoto(ctx context.Context) ([]_entities.Photo, error) {
	photos, err := ps.photoRepository.GetAllPhoto(ctx)
	return photos, err
}

func (ps *PhotoService) UpdatePhoto(ctx context.Context, updatePhoto _entities.Photo, idPhoto int, idToken int) (_entities.Photo, error) {
	getPhoto, err := ps.photoRepository.GetPhoto(ctx, idPhoto)
	if err != nil {
		return getPhoto, err
	}

	//validasi user login
	if idToken != getPhoto.User_id {
		return getPhoto, errors.New("unauthorized")
	}

	//validasi update photo
	if updatePhoto.Title != "" {
		getPhoto.Title = updatePhoto.Title
	}
	if updatePhoto.Caption != "" {
		getPhoto.Caption = updatePhoto.Caption
	}
	if updatePhoto.Photo_url != "" {
		getPhoto.Photo_url = updatePhoto.Photo_url

	}

	photo, err := ps.photoRepository.UpdatePhoto(ctx, getPhoto, idPhoto)
	photo.Id = idPhoto
	photo.Updated_at.Time = time.Now()
	return photo, err
}

func (ps *PhotoService) DeletePhoto(ctx context.Context, idPhoto int, idToken int) error {
	photo, errGetPhoto := ps.photoRepository.GetPhoto(ctx, idPhoto)
	if errGetPhoto != nil {
		return errGetPhoto
	}

	if idToken != photo.User_id {
		return errors.New("unauthorized")
	}

	err := ps.photoRepository.DeletePhoto(ctx, idPhoto)
	return err
}
