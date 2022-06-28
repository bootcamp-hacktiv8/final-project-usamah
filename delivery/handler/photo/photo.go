package photo

import (
	"database/sql"
	"encoding/json"
	_helper "final-project-usamah/delivery/helper"
	_response "final-project-usamah/delivery/helper/response"
	"final-project-usamah/delivery/middlewares"
	_entities "final-project-usamah/entities"
	_photoService "final-project-usamah/service/photo"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PhotoHandler struct {
	photoService _photoService.PhotoServiceInterface
}

func NewPhotoHandler(photoService _photoService.PhotoServiceInterface) *PhotoHandler {
	return &PhotoHandler{
		photoService: photoService,
	}
}

func (ph *PhotoHandler) CreatePhotoHandler(w http.ResponseWriter, r *http.Request) {
	//proses mengambil data user login
	ctx := r.Context()
	userLogin := middlewares.ForContext(ctx)

	//proses request body dari user
	var newPhoto _entities.Photo
	errDecode := json.NewDecoder(r.Body).Decode(&newPhoto)
	if errDecode != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	photo, err := ph.photoService.CreatePhoto(newPhoto, userLogin.Id)

	photoResponse := _response.ResponsePhoto(photo)

	if err != nil {
		response, _ := json.Marshal(_helper.APIResponseFailed(err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	} else {
		response, _ := json.Marshal(_helper.APIResponseSuccess("success create photo", photoResponse))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (ph *PhotoHandler) GetAllPhotoHandler(w http.ResponseWriter, r *http.Request) {
	photos, err := ph.photoService.GetAllPhoto()
	responsePhotos := _response.ResponseGetPhoto(photos)
	switch {
	case err == sql.ErrNoRows:
		response, _ := json.Marshal(_helper.APIResponseFailed("data not found"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(response)
	case err != nil:
		response, _ := json.Marshal(_helper.APIResponseFailed(err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		response, _ := json.Marshal(_helper.APIResponseSuccess("success get all photo", responsePhotos))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (ph *PhotoHandler) UpdatePhotoHandler(w http.ResponseWriter, r *http.Request) {
	//proses mengambil id param
	params := mux.Vars(r)
	idStr := params["photoId"]
	id, _ := strconv.Atoi(idStr)

	//proses mengambil data user login
	ctx := r.Context()
	userLogin := middlewares.ForContext(ctx)

	//proses request body dari user
	var updatePhoto _entities.Photo
	errDecode := json.NewDecoder(r.Body).Decode(&updatePhoto)
	if errDecode != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	photo, err := ph.photoService.UpdatePhoto(updatePhoto, id, userLogin.Id)
	photoResponse := _response.ResponseUpdatePhoto(photo)
	switch {
	case err == sql.ErrNoRows: //check data is null?
		response, _ := json.Marshal(_helper.APIResponseFailed("data not found"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(response)
	case err != nil:
		response, _ := json.Marshal(_helper.APIResponseFailed(err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		response, _ := json.Marshal(_helper.APIResponseSuccess("success update photo", photoResponse))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (ph *PhotoHandler) DeletePhotoHandler(w http.ResponseWriter, r *http.Request) {
	//proses mengambil id param
	params := mux.Vars(r)
	idStr := params["photoId"]
	id, _ := strconv.Atoi(idStr)

	//proses mengambil data user login
	ctx := r.Context()
	userLogin := middlewares.ForContext(ctx)

	err := ph.photoService.DeletePhoto(id, userLogin.Id)
	switch {
	case err == sql.ErrNoRows: //check data is null?
		response, _ := json.Marshal(_helper.APIResponseFailed("data not found"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(response)
	case err != nil:
		response, _ := json.Marshal(_helper.APIResponseFailed(err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		response, _ := json.Marshal(_helper.APIResponseSuccessWithouData("your photo has been successfully deleted"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
