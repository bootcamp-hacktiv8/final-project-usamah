package social_media

import (
	"context"
	_entities "final-project-usamah/entities"
)

type SocmedServiceInterface interface {
	CreateSocmed(ctx context.Context, newSocmed _entities.Social_media, idToken int) (_entities.Social_media, error)
	GetAllSocmed(ctx context.Context) ([]_entities.Social_media, error)
	UpdateSocmed(ctx context.Context, updateSocmed _entities.Social_media, idSocmed int, idToken int) (_entities.Social_media, error)
	DeleteSocmed(ctx context.Context, idSocmed int, idToken int) error
}
