package comment

import (
	response "final-project-usamah/delivery/helper/response/comment"
	_entities "final-project-usamah/entities"
)

type CommentRepositoryInterface interface {
	CreateComment(newComment _entities.Comment) (_entities.Comment, int, error)
	GetAllComment() ([]response.FormatGetComment, error)
	GetComment(idComment int) (_entities.Comment, error)
	UpdateComment(updateComment _entities.Comment, idComment int) (_entities.Comment, error)
	DeleteComment(idComment int) error
}
