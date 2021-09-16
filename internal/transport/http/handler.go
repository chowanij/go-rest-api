package http

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

// Handler - stores pointer to comand service
type Handler struct {
	Router *mux.Router
}

// NewHandler - returns pointer to a handler
func NewHandler() *Handler {
	return &Handler{}
}

//SetupRoutes - sets all routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive")
	})
}
