package user

import (
	"encoding/json"
	_helper "final-project-usamah/delivery/helper"
	_response "final-project-usamah/delivery/helper/response"
	"final-project-usamah/delivery/middlewares"
	_entities "final-project-usamah/entities"
	_userService "final-project-usamah/service/user"
	"net/http"
)

type UserHandler struct {
	userService _userService.UserServiceInterface
}

func NewUserHandler(userService _userService.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	//proses request body dari user
	var newUser _entities.User
	errDecode := json.NewDecoder(r.Body).Decode(&newUser)
	if errDecode != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	user, err := uh.userService.Register(newUser)

	userResponse := _response.ResponseUser(user)

	if err != nil {
		response, _ := json.Marshal(_helper.APIResponseFailed(err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	} else {
		response, _ := json.Marshal(_helper.APIResponseSuccess("success create user", userResponse))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (uh *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	//proses request body dari user
	var inputLogin _helper.LoginInput

	errDecode := json.NewDecoder(r.Body).Decode(&inputLogin)
	if errDecode != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	user, errLogin := uh.userService.Login(inputLogin)
	if errLogin != nil {
		response, _ := json.Marshal(_helper.APIResponseFailed(errLogin.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	token, errToken := middlewares.GenerateToken(user.Id)
	if errToken != nil {
		response, _ := json.Marshal(_helper.APIResponseFailed("error generate token"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	responseLogin := _response.ResponseLogin(token)
	response, _ := json.Marshal(_helper.APIResponseSuccess("success login", responseLogin))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (uh *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	//proses mengambil data user login
	ctx := r.Context()
	userLogin := middlewares.ForContext(ctx)

	//proses request body dari user
	var updateUser _entities.User
	errDecode := json.NewDecoder(r.Body).Decode(&updateUser)
	if errDecode != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	user, err := uh.userService.UpdateUser(updateUser, userLogin.Id)
	userResponse := _response.ResponseUpdateUser(user)
	switch {
	case err != nil:
		response, _ := json.Marshal(_helper.APIResponseFailed(err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		response, _ := json.Marshal(_helper.APIResponseSuccess("success update user", userResponse))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (uh *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	//proses mengambil data user login
	ctx := r.Context()
	userLogin := middlewares.ForContext(ctx)

	err := uh.userService.DeleteUser(userLogin.Id)
	switch {
	case err != nil:
		response, _ := json.Marshal(_helper.APIResponseFailed(err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		response, _ := json.Marshal(_helper.APIResponseSuccessWithouData("your account has been successfully deleted"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
