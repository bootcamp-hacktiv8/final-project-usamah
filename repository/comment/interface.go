package comment

import (
	"context"
	_entities "final-project-usamah/entities"
)

type CommentRepositoryInterface interface {
	CreateComment(ctx context.Context, newComment _entities.Comment) (_entities.Comment, int, error)
	GetAllComment(ctx context.Context) ([]_entities.Comment, error)
	GetComment(ctx context.Context, idComment int) (_entities.Comment, error)
	UpdateComment(ctx context.Context, updateComment _entities.Comment, idComment int) (_entities.Comment, error)
	DeleteComment(ctx context.Context, idComment int) error
}
