package comment

import (
	"database/sql"
	"encoding/json"
	_helper "final-project-usamah/delivery/helper"
	_response "final-project-usamah/delivery/helper/response"
	"final-project-usamah/delivery/middlewares"
	_entities "final-project-usamah/entities"
	_commentService "final-project-usamah/service/comment"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CommentHandler struct {
	commentService _commentService.CommentServiceInterface
}

func NewCommentHandler(commentService _commentService.CommentServiceInterface) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

func (ch *CommentHandler) CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	//proses mengambil data user login
	ctx := r.Context()
	userLogin := middlewares.ForContext(ctx)

	//proses request body dari user
	var newComment _entities.Comment
	errDecode := json.NewDecoder(r.Body).Decode(&newComment)
	if errDecode != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	comment, err := ch.commentService.CreateComment(ctx, newComment, userLogin.Id)

	commentResponse := _response.ResponseComment(comment)

	if err != nil {
		response, _ := json.Marshal(_helper.APIResponseFailed(err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	} else {
		response, _ := json.Marshal(_helper.APIResponseSuccess("success create comment", commentResponse))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (ch *CommentHandler) GetAllCommentHandler(w http.ResponseWriter, r *http.Request) {
	comments, err := ch.commentService.GetAllComment(r.Context())
	responseComments := _response.ResponseGetComment(comments)
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
		response, _ := json.Marshal(_helper.APIResponseSuccess("success get all comment", responseComments))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (ch *CommentHandler) UpdateCommentHandler(w http.ResponseWriter, r *http.Request) {
	//proses mengambil id param
	params := mux.Vars(r)
	idStr := params["commentId"]
	id, _ := strconv.Atoi(idStr)

	//proses mengambil data user login
	ctx := r.Context()
	userLogin := middlewares.ForContext(ctx)

	//proses request body dari user
	var updateComment _entities.Comment
	errDecode := json.NewDecoder(r.Body).Decode(&updateComment)
	if errDecode != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	comment, err := ch.commentService.UpdateComment(ctx, updateComment, id, userLogin.Id)
	commentResponse := _response.ResponseUpdateComment(comment)
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
		response, _ := json.Marshal(_helper.APIResponseSuccess("success update comment", commentResponse))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (ch *CommentHandler) DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	//proses mengambil id param
	params := mux.Vars(r)
	idStr := params["commentId"]
	id, _ := strconv.Atoi(idStr)

	//proses mengambil data user login
	ctx := r.Context()
	userLogin := middlewares.ForContext(ctx)

	err := ch.commentService.DeleteComment(ctx, id, userLogin.Id)
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
		response, _ := json.Marshal(_helper.APIResponseSuccessWithouData("your comment has been successfully deleted"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
