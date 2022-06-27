package comment

import (
	"context"
	"database/sql"
	_entities "final-project-usamah/entities"
	"time"

	_ "github.com/lib/pq"
)

type CommentRepository struct {
	database *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{
		database: db,
	}
}

func (cr *CommentRepository) CreateComment(newComment _entities.Comment) (_entities.Comment, int, error) {
	query := "INSERT INTO comments (user_id, photo_id, message, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	ctx := context.Background()
	var id int
	err := cr.database.QueryRowContext(ctx, query, newComment.User_id, newComment.Photo_id, newComment.Message, newComment.Created_at).Scan(&id)
	if err != nil {
		return newComment, id, err
	}
	return newComment, id, nil
}

func (cr *CommentRepository) GetAllComment() ([]_entities.Comment, error) {
	query := `SELECT comments.id, comments.user_id, comments.photo_id, comments.message, comments.created_at, comments.updated_at,
	users.id, users.email, users.username, photos.id, photos.user_id, photos.title, photos.caption, photos.photo_url
	FROM comments
	JOIN users ON (comments.user_id = users.id)
	JOIN photos ON (comments.photo_id = photos.id)`

	ctx := context.Background()

	rows, err := cr.database.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []_entities.Comment
	for rows.Next() {
		var comment _entities.Comment
		err := rows.Scan(&comment.Id, &comment.User_id, &comment.Photo_id, &comment.Message, &comment.Created_at, &comment.Updated_at, &comment.User.Id, &comment.User.Email, &comment.User.Username, &comment.Photo.Id, &comment.Photo.User_id, &comment.Photo.Title, &comment.Photo.Caption, &comment.Photo.Photo_url)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (cr *CommentRepository) GetComment(idComment int) (_entities.Comment, error) {
	query := `SELECT id, user_id, photo_id, message, created_at, updated_at
	FROM comments WHERE id = $1`
	ctx := context.Background()
	var comment _entities.Comment

	err := cr.database.QueryRowContext(ctx, query, idComment).Scan(&comment.Id, &comment.User_id, &comment.Photo_id, &comment.Message, &comment.Created_at, &comment.Updated_at)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (cr *CommentRepository) UpdateComment(updateComment _entities.Comment, idComment int) (_entities.Comment, error) {
	query := `UPDATE photos SET message = 1$, updated_at = $2
	WHERE id = $3`
	ctx := context.Background()

	_, err := cr.database.ExecContext(ctx, query, updateComment.Message, time.Now(), idComment)
	if err != nil {
		return updateComment, err
	}
	return updateComment, nil
}

func (cr *CommentRepository) DeleteComment(idComment int) error {
	query := `DELETE FROM comments WHERE id = $1`
	ctx := context.Background()

	_, err := cr.database.ExecContext(ctx, query, idComment)
	if err != nil {
		return err
	}
	return nil
}
