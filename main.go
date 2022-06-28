package main

import (
	_configs "final-project-usamah/configs"
	"final-project-usamah/delivery/middlewares"
	"final-project-usamah/delivery/routes"

	_userHandler "final-project-usamah/delivery/handler/user"
	_userRepository "final-project-usamah/repository/user"
	_userService "final-project-usamah/service/user"

	_photoHandler "final-project-usamah/delivery/handler/photo"
	_photoRepository "final-project-usamah/repository/photo"
	_photoService "final-project-usamah/service/photo"

	_commentHandler "final-project-usamah/delivery/handler/comment"
	_commentRepository "final-project-usamah/repository/comment"
	_commentService "final-project-usamah/service/comment"

	_sosmedHandler "final-project-usamah/delivery/handler/social_media"
	_sosmedRepository "final-project-usamah/repository/social_media"
	_sosmedService "final-project-usamah/service/social_media"

	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db := _configs.GetConnection()
	defer db.Close()
	fmt.Println("Successfully connected to database")

	userRepository := _userRepository.NewUserRepository(db)
	userService := _userService.NewUserService(userRepository)
	userHandler := _userHandler.NewUserHandler(userService)

	photoRepository := _photoRepository.NewPhotoRepository(db)
	photoService := _photoService.NewPhotoService(photoRepository)
	photoHandler := _photoHandler.NewPhotoHandler(photoService)

	commentRepository := _commentRepository.NewCommentRepository(db)
	commentService := _commentService.NewCommentService(commentRepository)
	commentHandler := _commentHandler.NewCommentHandler(commentService)

	sosmedRepository := _sosmedRepository.NewSosmedRepository(db)
	sosmedService := _sosmedService.NewSosmedService(sosmedRepository)
	sosmedHandler := _sosmedHandler.NewSosmedHandler(sosmedService)

	r := mux.NewRouter()
	r.Use(middlewares.LoggingMiddleware)

	routes.UserAuthPath(r, userHandler)
	routes.PhotoPath(r, photoHandler)
	routes.CommentPath(r, commentHandler)
	routes.SocialMediaPath(r, sosmedHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
