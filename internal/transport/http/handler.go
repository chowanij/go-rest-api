package http

import (
	"encoding/json"
	"errors"
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

// BasicAuth - a handy middleware function that logs out incoming requests
func BasicAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if user == "admin" && pass == "password" && ok {
			original(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
		}
	}
}

// LoggingMiddleware - middleware to wrap request with logging
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"Method": r.Method,
				"Path": r.URL.Path,
			},
		).Infof("Handled request")
		log.Info()
		next.ServeHTTP(w, r)
	})
}

// SetupRoutes - sets all routes for our application
func (h *Handler) SetupRoutes() {
	log.Info("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")

	h.Router.HandleFunc("/api/comment", BasicAuth(h.PostComment)).Methods("POST")

	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")

	h.Router.HandleFunc("/api/comment/{id}", BasicAuth(h.UpdateComment)).Methods("PUT")

	h.Router.HandleFunc("/api/comment/{id}", BasicAuth(h.DeleteComment)).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		if err := sendOkResponse(w, Response{Message: "I am alive"}); err != nil {
			panic(err)
		}
	})
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
