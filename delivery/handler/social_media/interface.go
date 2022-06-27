package social_media

import "net/http"

type SosmedHandlerInterface interface {
	CreateSosmedHandler(w http.ResponseWriter, r *http.Request)
	GetAllSosmedHandler(w http.ResponseWriter, r *http.Request)
	UpdateSosmedHandler(w http.ResponseWriter, r *http.Request)
	DeleteSosmedHandler(w http.ResponseWriter, r *http.Request)
}
