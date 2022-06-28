package social_media

import (
	"context"
	_entities "final-project-usamah/entities"
)

type SocmedRepositoryInterface interface {
	CreateSocmed(ctx context.Context, newSocmed _entities.Social_media) (_entities.Social_media, int, error)
	GetAllSocmed(ctx context.Context) ([]_entities.Social_media, error)
	GetSocmed(ctx context.Context, idSocmed int) (_entities.Social_media, error)
	UpdateSocmed(ctx context.Context, updateSocmed _entities.Social_media, idSocmed int) (_entities.Social_media, error)
	DeleteSocmed(ctx context.Context, idSocmed int) error
}
