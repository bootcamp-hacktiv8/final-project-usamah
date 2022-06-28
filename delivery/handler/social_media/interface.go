package social_media

import "net/http"

type SocmedHandlerInterface interface {
	CreateSocmedHandler(w http.ResponseWriter, r *http.Request)
	GetAllSocmedHandler(w http.ResponseWriter, r *http.Request)
	UpdateSocmedHandler(w http.ResponseWriter, r *http.Request)
	DeleteSocmedHandler(w http.ResponseWriter, r *http.Request)
}
