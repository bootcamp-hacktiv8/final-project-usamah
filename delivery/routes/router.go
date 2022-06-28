package routes

import (
	_commentHandler "final-project-usamah/delivery/handler/comment"
	_photoHandler "final-project-usamah/delivery/handler/photo"
	_sosmedHandler "final-project-usamah/delivery/handler/social_media"
	_userHandler "final-project-usamah/delivery/handler/user"
	"final-project-usamah/delivery/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func UserAuthPath(r *mux.Router, uh *_userHandler.UserHandler) {
	r.HandleFunc("/users/register", uh.RegisterHandler).Methods("POST")
	r.HandleFunc("/users/login", uh.LoginHandler).Methods("POST")
	r.Handle("/users", middlewares.Authentication(http.HandlerFunc(uh.UpdateUserHandler))).Methods("PUT")
	r.Handle("/users", middlewares.Authentication(http.HandlerFunc(uh.DeleteUserHandler))).Methods("DELETE")
}

func PhotoPath(r *mux.Router, ph *_photoHandler.PhotoHandler) {
	r.Handle("/photos", middlewares.Authentication(http.HandlerFunc(ph.CreatePhotoHandler))).Methods("POST")
	r.Handle("/photos", middlewares.Authentication(http.HandlerFunc(ph.GetAllPhotoHandler))).Methods("GET")
	r.Handle("/photos/{photoId}", middlewares.Authentication(http.HandlerFunc(ph.UpdatePhotoHandler))).Methods("PUT")
	r.Handle("/photos/{photoId}", middlewares.Authentication(http.HandlerFunc(ph.DeletePhotoHandler))).Methods("DELETE")
}

func CommentPath(r *mux.Router, ch *_commentHandler.CommentHandler) {
	r.Handle("/comments", middlewares.Authentication(http.HandlerFunc(ch.CreateCommentHandler))).Methods("POST")
	r.Handle("/comments", middlewares.Authentication(http.HandlerFunc(ch.GetAllCommentHandler))).Methods("GET")
	r.Handle("/comments/{commentId}", middlewares.Authentication(http.HandlerFunc(ch.UpdateCommentHandler))).Methods("PUT")
	r.Handle("/comments/{commentId}", middlewares.Authentication(http.HandlerFunc(ch.DeleteCommentHandler))).Methods("DELETE")
}

func SocialMediaPath(r *mux.Router, sh *_sosmedHandler.SosmedHandler) {
	r.Handle("/socialmedias", middlewares.Authentication(http.HandlerFunc(sh.CreateSosmedHandler))).Methods("POST")
	r.Handle("/socialmedias", middlewares.Authentication(http.HandlerFunc(sh.GetAllSosmedHandler))).Methods("GET")
	r.Handle("/socialmedias/{socialMediaId}", middlewares.Authentication(http.HandlerFunc(sh.UpdateSosmedHandler))).Methods("PUT")
	r.Handle("/socialmedias/{socialMediaId}", middlewares.Authentication(http.HandlerFunc(sh.DeleteSosmedHandler))).Methods("DELETE")
}
