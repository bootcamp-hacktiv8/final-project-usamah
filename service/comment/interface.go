package comment

import (
	_entities "final-project-usamah/entities"
)

type CommentServiceInterface interface {
	CreateComment(newComment _entities.Comment, idToken int) (_entities.Comment, error)
	GetAllComment() ([]_entities.Comment, error)
	UpdateComment(updateComment _entities.Comment, idComment int, idToken int) (_entities.Comment, error)
	DeleteComment(idComment int, idToken int) error
}
