package routes

import (
	_commentRouter "final-project-usamah/delivery/routes/comment"
	_photoRouter "final-project-usamah/delivery/routes/photo"
	_sosmedRouter "final-project-usamah/delivery/routes/social_media"
	_userRouter "final-project-usamah/delivery/routes/user"
	_commentRepository "final-project-usamah/repository/comment"
	_photoRepository "final-project-usamah/repository/photo"
	_sosmedRepository "final-project-usamah/repository/social_media"
	_userRepository "final-project-usamah/repository/user"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Routes(
	userRepository _userRepository.UserRepositoryInterface,
	photoRepository _photoRepository.PhotoRepositoryInterface,
	commentRepository _commentRepository.CommentRepositoryInterface,
	sosmedRepository _sosmedRepository.SosmedRepositoryInterface,
) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	mount(router, "/users", _userRouter.UserResource{}.UserRoute(userRepository))
	mount(router, "/photos", _photoRouter.PhotoResource{}.PhotoRoute(photoRepository))
	mount(router, "/comments", _commentRouter.CommentResource{}.CommentRoute(commentRepository))
	mount(router, "/socialmedias", _sosmedRouter.SosmedResource{}.SosmedRoute(sosmedRepository))

	return router
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
