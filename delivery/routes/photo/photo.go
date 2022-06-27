package photo

import (
	_photoHandler "final-project-usamah/delivery/handler/photo"
	"final-project-usamah/delivery/middlewares"
	_photoRepository "final-project-usamah/repository/photo"
	_photoService "final-project-usamah/service/photo"
	"net/http"

	"github.com/gorilla/mux"
)

type PhotoResource struct{}

func (pr PhotoResource) PhotoRoute(photoRepository _photoRepository.PhotoRepositoryInterface) *mux.Router {
	photoService := _photoService.NewPhotoService(photoRepository)
	photoHandler := _photoHandler.NewPhotoHandler(photoService)

	router := mux.NewRouter()
	router.Handle("/", middlewares.Authentication(http.HandlerFunc(photoHandler.CreatePhotoHandler))).Methods("POST")
	router.Handle("/", middlewares.Authentication(http.HandlerFunc(photoHandler.GetAllPhotoHandler))).Methods("GET")
	router.Handle("/{photoId}", middlewares.Authentication(http.HandlerFunc(photoHandler.UpdatePhotoHandler))).Methods("PUT")
	router.Handle("/{photoId}", middlewares.Authentication(http.HandlerFunc(photoHandler.DeletePhotoHandler))).Methods("DELETE")
	return router
}
