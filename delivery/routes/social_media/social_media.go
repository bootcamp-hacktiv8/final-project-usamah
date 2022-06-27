package social_media

import (
	_sosmedHandler "final-project-usamah/delivery/handler/social_media"
	"final-project-usamah/delivery/middlewares"
	_sosmedRepository "final-project-usamah/repository/social_media"
	_sosmedService "final-project-usamah/service/social_media"
	"net/http"

	"github.com/gorilla/mux"
)

type SosmedResource struct{}

func (sr SosmedResource) SosmedRoute(sosmedRepository _sosmedRepository.SosmedRepositoryInterface) *mux.Router {
	sosmedService := _sosmedService.NewSosmedService(sosmedRepository)
	sosmedHandler := _sosmedHandler.NewSosmedHandler(sosmedService)

	router := mux.NewRouter()
	router.Handle("/", middlewares.Authentication(http.HandlerFunc(sosmedHandler.CreateSosmedHandler))).Methods("POST")
	router.Handle("/", middlewares.Authentication(http.HandlerFunc(sosmedHandler.GetAllSosmedHandler))).Methods("GET")
	router.Handle("/{socialMediaId}", middlewares.Authentication(http.HandlerFunc(sosmedHandler.UpdateSosmedHandler))).Methods("PUT")
	router.Handle("/{socialMediaId}", middlewares.Authentication(http.HandlerFunc(sosmedHandler.DeleteSosmedHandler))).Methods("DELETE")
	return router
}
