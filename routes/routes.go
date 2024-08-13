package routes

import (
	"github.com/gorilla/mux"
	"userauth/handlers"
	"userauth/middlewares"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	// Protected routes
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.Use(middlewares.JWTAuthMiddleware)
	userRouter.HandleFunc("", handlers.GetUsers).Methods("GET")
	userRouter.HandleFunc("/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	userRouter.HandleFunc("", handlers.CreateUser).Methods("POST")
	userRouter.HandleFunc("/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	return router
}
