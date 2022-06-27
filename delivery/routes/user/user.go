package user

import (
	_userHandler "final-project-usamah/delivery/handler/user"
	"final-project-usamah/delivery/middlewares"
	_userRepository "final-project-usamah/repository/user"
	_userService "final-project-usamah/service/user"
	"net/http"

	"github.com/gorilla/mux"
)

type UserResource struct{}

func (ur UserResource) UserRoute(userRepository _userRepository.UserRepositoryInterface) *mux.Router {
	userService := _userService.NewUserService(userRepository)
	userHandler := _userHandler.NewUserHandler(userService)

	router := mux.NewRouter()
	router.Handle("/register", http.HandlerFunc(userHandler.RegisterHandler)).Methods("POST")
	router.Handle("/login", http.HandlerFunc(userHandler.LoginHandler)).Methods("POST")
	router.Handle("/", middlewares.Authentication(http.HandlerFunc(userHandler.UpdateUserHandler))).Methods("PUT")
	router.Handle("/", middlewares.Authentication(http.HandlerFunc(userHandler.DeleteUserHandler))).Methods("DELETE")
	router.Use(middlewares.LoggingMiddleware)
	return router
}
