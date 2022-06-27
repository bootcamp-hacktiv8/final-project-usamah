package comment

import (
	_entities "final-project-usamah/entities"
)

type CommentRepositoryInterface interface {
	CreateComment(newComment _entities.Comment) (_entities.Comment, int, error)
	GetAllComment() ([]_entities.Comment, error)
	GetComment(idComment int) (_entities.Comment, error)
	UpdateComment(updateComment _entities.Comment, idComment int) (_entities.Comment, error)
	DeleteComment(idComment int) error
}
