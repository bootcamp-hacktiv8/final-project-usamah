package main

import (
	_configs "final-project-usamah/configs"
	_routes "final-project-usamah/delivery/routes"
	_commentRepository "final-project-usamah/repository/comment"
	_photoRepository "final-project-usamah/repository/photo"
	_sosmedRepository "final-project-usamah/repository/social_media"
	_userRepository "final-project-usamah/repository/user"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	db := _configs.GetConnection()
	defer db.Close()
	fmt.Println("Successfully connected to database")

	userRepository := _userRepository.NewUserRepository(db)
	photoRepository := _photoRepository.NewPhotoRepository(db)
	commentRepository := _commentRepository.NewCommentRepository(db)
	sosmedRepository := _sosmedRepository.NewSosmedRepository(db)

	router := _routes.Routes(
		userRepository,
		photoRepository,
		commentRepository,
		sosmedRepository,
	)

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
