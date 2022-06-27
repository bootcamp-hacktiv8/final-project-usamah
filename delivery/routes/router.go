package routes

import (
	_commentRouter "final-project-usamah/delivery/routes/comment"
	_photoRouter "final-project-usamah/delivery/routes/photo"
	_userRouter "final-project-usamah/delivery/routes/user"
	_commentRepository "final-project-usamah/repository/comment"
	_photoRepository "final-project-usamah/repository/photo"
	_userRepository "final-project-usamah/repository/user"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Routes(
	userRepository _userRepository.UserRepositoryInterface,
	photoRepository _photoRepository.PhotoRepositoryInterface,
	commentRepository _commentRepository.CommentRepositoryInterface,
) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	mount(router, "/users", _userRouter.UserResource{}.UserRoute(userRepository))
	mount(router, "/photos", _photoRouter.PhotoResource{}.PhotoRoute(photoRepository))
	mount(router, "/comments", _commentRouter.CommentResource{}.CommentRoute(commentRepository))

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
