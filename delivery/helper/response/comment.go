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

func ResponseComment(comment entities.Comment) CommentFormatter {
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

type GetCommentFormatter struct {
	Id         int             `json:"id"`
	User_id    int             `json:"user_id"`
	Photo_id   int             `json:"photo_id"`
	Message    string          `json:"message"`
	Created_at time.Time       `json:"created_at"`
	Updated_at sql.NullTime    `json:"updated_at"`
	User       GetCommentUser  `json:"user"`
	Photo      GetCommentPhoto `json:"photo"`
}

func ResponseGetComment(comment []entities.Comment) []GetCommentFormatter {
	var comments []GetCommentFormatter
	for i := 0; i < len(comment); i++ {
		getcomment := GetCommentFormatter{
			Id:         comment[i].Id,
			User_id:    comment[i].User_id,
			Photo_id:   comment[i].Photo_id,
			Message:    comment[i].Message,
			Created_at: comment[i].Created_at,
			Updated_at: comment[i].Updated_at,
			User: GetCommentUser{
				Id:       comment[i].User.Id,
				Username: comment[i].User.Username,
				Email:    comment[i].User.Email,
			},
			Photo: GetCommentPhoto{
				Id:        comment[i].Photo.Id,
				User_id:   comment[i].Photo.User_id,
				Title:     comment[i].Photo.Title,
				Caption:   comment[i].Photo.Caption,
				Photo_url: comment[i].Photo.Photo_url,
			},
		}
		comments = append(comments, getcomment)
	}
	return comments
}

type UpdateCommentFormatter struct {
	Id         int          `json:"id"`
	User_id    int          `json:"user_id"`
	Photo_id   int          `json:"photo_id"`
	Message    string       `json:"message"`
	Updated_at sql.NullTime `json:"updated_at"`
}

func ResponseUpdateComment(comment entities.Comment) UpdateCommentFormatter {
	formatter := UpdateCommentFormatter{
		Id:         comment.Id,
		User_id:    comment.User_id,
		Photo_id:   comment.Photo_id,
		Message:    comment.Message,
		Updated_at: comment.Updated_at,
	}
	return formatter
}
