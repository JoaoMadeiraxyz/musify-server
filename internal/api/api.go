package api

import (
	"log/slog"
	"net/http"
	"sync"

	"github.com/JoaoMadeiraxyz/musify-server/internal/store/pgstore"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type apiHandler struct {
	q  *pgstore.Queries
	r  *chi.Mux
	mu *sync.Mutex
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler{
		q:  q,
		mu: &sync.Mutex{},
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/", a.handleCreateUser)
			r.Get("/", a.handleGetUsers)
		})
	})

	a.r = r
	return a
}

func (h apiHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) {}
func (h apiHandler) handleGetUsers(w http.ResponseWriter, r *http.Request)   {
	users, err := h.q.GetUsers(r.Context())
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		slog.Error("failed to get users", "error", err)
		return
	}

	if users == nil{
		users = []pgstore.User{}
	}

	sendJSON(w, users)
}
