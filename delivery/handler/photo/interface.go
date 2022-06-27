package photo

import "net/http"

type PhotoHandlerInterface interface {
	CreatePhotoHandler(w http.ResponseWriter, r *http.Request)
	GetAllPhotoHandler(w http.ResponseWriter, r *http.Request)
	UpdatePhotoHandler(w http.ResponseWriter, r *http.Request)
	DeletePhotoHandler(w http.ResponseWriter, r *http.Request)
}
