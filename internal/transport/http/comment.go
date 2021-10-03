package http

import (
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/chowanij/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
)

// GetAllComments - retrieve all comments
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComents()

	if err != nil {
		sendErrorResponse(w, "Error receiving all comments", err)
		return
	}

	if err := sendOkResponse(w, comments); err != nil {
		panic(err)
	}

}

// GetComment - retrieve comment by id
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		sendErrorResponse(w, "Failed to convert ID to UINT", err)
		return
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Failed to retrieve comment by ID", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

// PostComment - post new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var c comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		sendErrorResponse(w, "Failed to post decode", err)
		return
	}

	comment, err := h.Service.PostComment(c)

	if err != nil {
		sendErrorResponse(w, "Failed to post comment", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}

}

// UpdateComment - update comment with ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var c comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		sendErrorResponse(w, "Failed to decode", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)

	comment, err := h.Service.UpdateComment(uint(i), c)

	if err != nil {
		sendErrorResponse(w, "Failed to update comment", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

// DeleteComment - retrieve comment by id
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Comment sucessfully deleted",
		Error:   "",
	}
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		sendErrorResponse(w, "Failed to convert ID to UINT", err)
		return
	}

	err = h.Service.DeleteComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Failed to delete comment by ID", err)
		return
	}

	if err := sendOkResponse(w, response); err != nil {
		panic(err)
	}
}
