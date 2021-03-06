package comment

import (
	"context"
	"errors"
	_entities "final-project-usamah/entities"
	_commentRepository "final-project-usamah/repository/comment"
	"time"
)

type CommentService struct {
	commentRepository _commentRepository.CommentRepositoryInterface
}

func NewCommentService(commentRepository _commentRepository.CommentRepositoryInterface) CommentServiceInterface {
	return &CommentService{
		commentRepository: commentRepository,
	}
}

func (cs *CommentService) CreateComment(ctx context.Context, newComment _entities.Comment, idToken int) (_entities.Comment, error) {
	//validasi saat create photo
	if newComment.Message == "" {
		return newComment, errors.New("message is required")
	}

	newComment.User_id = idToken
	newComment.Created_at = time.Now()
	comment, id, err := cs.commentRepository.CreateComment(ctx, newComment)
	comment.Id = id
	return comment, err
}

func (cs *CommentService) GetAllComment(ctx context.Context) ([]_entities.Comment, error) {
	comments, err := cs.commentRepository.GetAllComment(ctx)
	return comments, err
}

func (cs *CommentService) UpdateComment(ctx context.Context, updateComment _entities.Comment, idComment int, idToken int) (_entities.Comment, error) {
	getComment, err := cs.commentRepository.GetComment(ctx, idComment)
	if err != nil {
		return getComment, err
	}

	//validasi user login
	if idToken != getComment.User_id {
		return getComment, errors.New("unauthorized")
	}

	//validasi update photo
	if updateComment.Message != "" {
		getComment.Message = updateComment.Message
	}

	comment, err := cs.commentRepository.UpdateComment(ctx, getComment, idComment)
	comment.Id = idComment
	comment.Updated_at.Time = time.Now()
	return comment, err
}

func (cs *CommentService) DeleteComment(ctx context.Context, idComment int, idToken int) error {
	comment, errGetComment := cs.commentRepository.GetComment(ctx, idComment)
	if errGetComment != nil {
		return errGetComment
	}

	if idToken != comment.User_id {
		return errors.New("unauthorized")
	}

	err := cs.commentRepository.DeleteComment(ctx, idComment)
	return err
}
