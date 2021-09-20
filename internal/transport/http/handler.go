package http

import (
	"encoding/json"
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

// Response - an object to store responses from our API
type Response struct {
	Message string
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

	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")

	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")

	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")

	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am alive"}); err != nil {
			panic(err)
		}
	})
}

// GetAllComments - retrieve all comments
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	comments, err := h.Service.GetAllComents()

	if err != nil {
		fmt.Fprintf(w, "Error receiving all comments")
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}

}

// GetComment - retrieve comment by id
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		fmt.Fprintf(w, "Failed to convert ID to UINT")
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		if err := json.NewEncoder(w).Encode(Response{Message: "Failed to retreive comment by id"}); err != nil {
			panic(err)
		}
	} else if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

// PostComment - post new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var c comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		if err := json.NewEncoder(w).Encode(Response{Message: "Failed to decode"}); err != nil {
			panic(err)
		}
	}

	comment, err := h.Service.PostComment(c)

	if err != nil {
		if err := json.NewEncoder(w).Encode(Response{Message: "Failed to post comment"}); err != nil {
			panic(err)
		}
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}

}

// UpdateComment - update comment with ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var c comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		if err := json.NewEncoder(w).Encode(Response{Message: "Failed to decode"}); err != nil {
			panic(err)
		}
	}

	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)

	comment, err := h.Service.UpdateComment(uint(i), c)

	if err != nil {
		fmt.Fprintf(w, "Failed to update comment")
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response{Message: "Comment sucessfully deleted"}); err != nil {
		panic(err)
	}
}
