package response

import (
	"database/sql"
	"final-project-usamah/entities"
	"time"
)

type CommentFormatter struct {
	Id         int       `json:"id"`
	User_id    int       `json:"user_id"`
	Photo_id   int       `json:"photo_id"`
	Message    string    `json:"message"`
	Created_at time.Time `json:"created_at"`
}

func FormatComment(comment entities.Comment) CommentFormatter {
	formatter := CommentFormatter{
		Id:         comment.Id,
		User_id:    comment.User_id,
		Photo_id:   comment.Photo_id,
		Message:    comment.Message,
		Created_at: comment.Created_at,
	}
	return formatter
}

type GetCommentUser struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type GetCommentPhoto struct {
	Id        int    `json:"id"`
	User_id   int    `json:"user_id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url"`
}

type FormatGetComment struct {
	Id         int             `json:"id"`
	User_id    int             `json:"user_id"`
	Photo_id   int             `json:"photo_id"`
	Message    string          `json:"message"`
	Created_at time.Time       `json:"created_at"`
	Updated_at sql.NullTime    `json:"updated_at"`
	User       GetCommentUser  `json:"user"`
	Photo      GetCommentPhoto `json:"photo"`
}

type UpdateComment struct {
	Id         int          `json:"id"`
	User_id    int          `json:"user_id"`
	Photo_id   int          `json:"photo_id"`
	Message    string       `json:"message"`
	Updated_at sql.NullTime `json:"updated_at"`
}

func FormatUpdateComment(comment entities.Comment) UpdateComment {
	formatter := UpdateComment{
		Id:         comment.Id,
		User_id:    comment.User_id,
		Photo_id:   comment.Photo_id,
		Message:    comment.Message,
		Updated_at: comment.Updated_at,
	}
	return formatter
}
