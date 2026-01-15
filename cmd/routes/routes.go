package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/r6rap/Quanta/internal/handler"
	"github.com/r6rap/Quanta/internal/middleware"
	"github.com/r6rap/Quanta/pkg/database"
	"github.com/r6rap/Quanta/pkg/response"
)

type Routes struct {
	db *database.DB
}

func NewRoutes(db *database.DB) *Routes {
	return &Routes{db: db}
}

func(r *Routes) InitRoutes() *mux.Router {
	authHandler := handler.NewAuthHandler(r.db)
	authMiddleware := middleware.NewAuthMiddleware(authHandler)

	route := mux.NewRouter()

	auth := route.PathPrefix("/quanta/auth").Subrouter()
	auth.HandleFunc("/register", authHandler.Register).Methods("POST")
	auth.HandleFunc("/login", authHandler.Login).Methods("POST")
	auth.HandleFunc("/logout", authHandler.Logout).Methods("POST")
	auth.HandleFunc("/refresh", authHandler.RefreshAccessToken).Methods("POST")

	api := route.PathPrefix("/quanta/api").Subrouter()
	api.Use(authMiddleware.Middleware)
	api.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id")
		email := r.Context().Value("email")

		response.OK(w, http.StatusOK, userID, email)
	})

	return route
}