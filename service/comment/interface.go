package comment

import (
	"context"
	_entities "final-project-usamah/entities"
)

type CommentServiceInterface interface {
	CreateComment(ctx context.Context, newComment _entities.Comment, idToken int) (_entities.Comment, error)
	GetAllComment(ctx context.Context) ([]_entities.Comment, error)
	UpdateComment(ctx context.Context, updateComment _entities.Comment, idComment int, idToken int) (_entities.Comment, error)
	DeleteComment(ctx context.Context, idComment int, idToken int) error
}
