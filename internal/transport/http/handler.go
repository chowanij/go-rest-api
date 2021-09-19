package http

import (
	"fmt"
	"strconv"

	"net/http"

	"github.com/chowanij/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
)

// Handler - stores pointer to comand service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// NewHandler - returns pointer to a handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - sets all routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")

	h.Router.HandleFunc("/api/comment{id}", h.PostComment).Methods("POST")

	h.Router.HandleFunc("/api/comment{id}", h.GetComment).Methods("GET")

	h.Router.HandleFunc("/api/comment{id}", h.UpdateComment).Methods("PUT")

	h.Router.HandleFunc("/api/comment{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive")
	})
}

// GetAllComments - retrieve all comments
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComents()

	if err != nil {
		fmt.Fprintf(w, "Error receiving all comments")
	}

	fmt.Fprintf(w, "%+v", comments)

}

// GetComment - retrieve comment by id
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		fmt.Fprintf(w, "Failed to convert ID to UINT")
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		fmt.Fprintf(w, "Failed receiving comment by ID")
	}

	fmt.Fprintf(w, "%+v", comment)
}

// PostComment - post new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.PostComment(comment.Comment{
		Slug: "/",
	})

	if err != nil {
		fmt.Fprintf(w, "Failed to post comment")
	}

	fmt.Fprintf(w, "%+v", comment)

}

// UpdateComment - update comment with ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.UpdateComment(1, comment.Comment{
		Slug: "/new",
	})

	if err != nil {
		fmt.Fprintf(w, "Failed to update comment")
	}

	fmt.Fprintf(w, "%+v", comment)
}

// DeleteComment - retrieve comment by id
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		fmt.Fprintf(w, "Failed to convert ID to UINT")
	}

	err = h.Service.DeleteComment(uint(i))
	if err != nil {
		fmt.Fprintf(w, "Failed to delete comment by ID")
	}

	fmt.Fprintf(w, "Successfully deleted comment")
}
