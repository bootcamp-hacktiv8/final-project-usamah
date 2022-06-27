package comment

import (
	_commentHandler "final-project-usamah/delivery/handler/comment"
	"final-project-usamah/delivery/middlewares"
	_commentRepository "final-project-usamah/repository/comment"
	_commentService "final-project-usamah/service/comment"
	"net/http"

	"github.com/gorilla/mux"
)

type CommentResource struct{}

func (cr CommentResource) CommentRoute(commentRepository _commentRepository.CommentRepositoryInterface) *mux.Router {
	commentService := _commentService.NewCommentService(commentRepository)
	commentHandler := _commentHandler.NewCommentHandler(commentService)

	router := mux.NewRouter()
	router.Handle("/", middlewares.Authentication(http.HandlerFunc(commentHandler.CreateCommentHandler))).Methods("POST")
	router.Handle("/", middlewares.Authentication(http.HandlerFunc(commentHandler.GetAllCommentHandler))).Methods("GET")
	router.Handle("/{commentId}", middlewares.Authentication(http.HandlerFunc(commentHandler.UpdateCommentHandler))).Methods("PUT")
	router.Handle("/{commentId}", middlewares.Authentication(http.HandlerFunc(commentHandler.DeleteCommentHandler))).Methods("DELETE")
	router.Use(middlewares.LoggingMiddleware)
	return router
}
