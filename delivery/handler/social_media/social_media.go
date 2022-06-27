package social_media

import (
	"database/sql"
	"encoding/json"
	_helper "final-project-usamah/delivery/helper"
	_response "final-project-usamah/delivery/helper/response/social_media"
	"final-project-usamah/delivery/middlewares"
	_entities "final-project-usamah/entities"
	_sosmedService "final-project-usamah/service/social_media"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SosmedHandler struct {
	sosmedService _sosmedService.SosmedServiceInterface
}

func NewSosmedHandler(sosmedService _sosmedService.SosmedServiceInterface) *SosmedHandler {
	return &SosmedHandler{
		sosmedService: sosmedService,
	}
}

func (sh *SosmedHandler) CreateSosmedHandler(w http.ResponseWriter, r *http.Request) {
	//proses mengambil data user login
	ctx := r.Context()
	userLogin := middlewares.ForContext(ctx)

	//proses request body dari user
	var newSosmed _entities.Social_media
	errDecode := json.NewDecoder(r.Body).Decode(&newSosmed)
	if errDecode != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	sosmed, err := sh.sosmedService.CreateSosmed(newSosmed, userLogin.Id)

	sosmedResponse := _response.FormatSocialMedia(sosmed)

	if err != nil {
		response, _ := json.Marshal(_helper.APIResponseFailed(err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	} else {
		response, _ := json.Marshal(_helper.APIResponseSuccess("success create social_media", sosmedResponse))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (sh *SosmedHandler) GetAllSosmedHandler(w http.ResponseWriter, r *http.Request) {
	sosmeds, err := sh.sosmedService.GetAllSosmed()
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
		response, _ := json.Marshal(_helper.APIResponseSuccess("success get all social_media", sosmeds))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (sh *SosmedHandler) UpdateSosmedHandler(w http.ResponseWriter, r *http.Request) {
	//proses mengambil id param
	params := mux.Vars(r)
	idStr := params["socialMediaId"]
	id, _ := strconv.Atoi(idStr)

	//proses mengambil data user login
	ctx := r.Context()
	userLogin := middlewares.ForContext(ctx)

	//proses request body dari user
	var updateSosmed _entities.Social_media
	errDecode := json.NewDecoder(r.Body).Decode(&updateSosmed)
	if errDecode != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	sosmed, err := sh.sosmedService.UpdateSosmed(updateSosmed, id, userLogin.Id)
	sosmedResponse := _response.FormatUpdateSosmed(sosmed)
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
		response, _ := json.Marshal(_helper.APIResponseSuccess("success update social_media", sosmedResponse))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (sh *SosmedHandler) DeleteSosmedHandler(w http.ResponseWriter, r *http.Request) {
	//proses mengambil id param
	params := mux.Vars(r)
	idStr := params["socialMediaId"]
	id, _ := strconv.Atoi(idStr)

	//proses mengambil data user login
	ctx := r.Context()
	userLogin := middlewares.ForContext(ctx)

	err := sh.sosmedService.DeleteSosmed(id, userLogin.Id)
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
		response, _ := json.Marshal(_helper.APIResponseSuccessWithouData("your social media has been successfully deleted"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
