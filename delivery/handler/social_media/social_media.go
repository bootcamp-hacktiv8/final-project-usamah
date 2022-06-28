package social_media

import (
	"database/sql"
	"encoding/json"
	_helper "final-project-usamah/delivery/helper"
	_response "final-project-usamah/delivery/helper/response"
	"final-project-usamah/delivery/middlewares"
	_entities "final-project-usamah/entities"
	_socmedService "final-project-usamah/service/social_media"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SocmedHandler struct {
	socmedService _socmedService.SocmedServiceInterface
}

func NewSocmedHandler(socmedService _socmedService.SocmedServiceInterface) *SocmedHandler {
	return &SocmedHandler{
		socmedService: socmedService,
	}
}

func (sh *SocmedHandler) CreateSocmedHandler(w http.ResponseWriter, r *http.Request) {
	//proses mengambil data user login
	ctx := r.Context()
	userLogin := middlewares.ForContext(ctx)

	//proses request body dari user
	var newSocmed _entities.Social_media
	errDecode := json.NewDecoder(r.Body).Decode(&newSocmed)
	if errDecode != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	socmed, err := sh.socmedService.CreateSocmed(newSocmed, userLogin.Id)

	socmedResponse := _response.ResponseSocialMedia(socmed)

	if err != nil {
		response, _ := json.Marshal(_helper.APIResponseFailed(err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	} else {
		response, _ := json.Marshal(_helper.APIResponseSuccess("success create social_media", socmedResponse))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (sh *SocmedHandler) GetAllSocmedHandler(w http.ResponseWriter, r *http.Request) {
	socmeds, err := sh.socmedService.GetAllSocmed()
	responseSocmeds := _response.ResponseGetSocialMedia(socmeds)
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
		response, _ := json.Marshal(_helper.APIResponseSuccess("success get all social_media", responseSocmeds))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (sh *SocmedHandler) UpdateSocmedHandler(w http.ResponseWriter, r *http.Request) {
	//proses mengambil id param
	params := mux.Vars(r)
	idStr := params["socialMediaId"]
	id, _ := strconv.Atoi(idStr)

	//proses mengambil data user login
	ctx := r.Context()
	userLogin := middlewares.ForContext(ctx)

	//proses request body dari user
	var updateSocmed _entities.Social_media
	errDecode := json.NewDecoder(r.Body).Decode(&updateSocmed)
	if errDecode != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	socmed, err := sh.socmedService.UpdateSocmed(updateSocmed, id, userLogin.Id)
	socmedResponse := _response.ResponseUpdateSocmed(socmed)
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
		response, _ := json.Marshal(_helper.APIResponseSuccess("success update social_media", socmedResponse))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (sh *SocmedHandler) DeleteSocmedHandler(w http.ResponseWriter, r *http.Request) {
	//proses mengambil id param
	params := mux.Vars(r)
	idStr := params["socialMediaId"]
	id, _ := strconv.Atoi(idStr)

	//proses mengambil data user login
	ctx := r.Context()
	userLogin := middlewares.ForContext(ctx)

	err := sh.socmedService.DeleteSocmed(id, userLogin.Id)
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
