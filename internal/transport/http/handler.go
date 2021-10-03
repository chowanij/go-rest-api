package http

import (
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/chowanij/go-rest-api/internal/comment"
	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

// Handler - stores pointer to comand service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// Response - an object to store responses from our API
type Response struct {
	Message string
	Error   string
}

// NewHandler - returns pointer to a handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - sets all routes for our application
func (h *Handler) SetupRoutes() {
	log.Info("Setting up routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")

	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")

	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")

	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")

	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		if err := sendOkResponse(w, Response{Message: "I am alive"}); err != nil {
			panic(err)
		}
	})
}

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

func sendOkResponse(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
