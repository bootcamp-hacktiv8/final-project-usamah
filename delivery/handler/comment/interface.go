package comment

import "net/http"

type CommentHandlerInterface interface {
	CreateCommentHandler(w http.ResponseWriter, r *http.Request)
	GetAllCommentHandler(w http.ResponseWriter, r *http.Request)
	UpdateCommentHandler(w http.ResponseWriter, r *http.Request)
	DeleteCommentHandler(w http.ResponseWriter, r *http.Request)
}
