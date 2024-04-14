package main

import (
	"net/http"
	"user-mgt/internal/database"
	"user-mgt/internal/delivery/http/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Initialize database with connection pool configurations
	db, err := database.InitDB(
		"postgresql://postgres:123456@localhost:5432/user-mgt",
		10,
		5,
	) // Example values: maxOpenConns = 10, maxIdleConns = 5
	if err != nil {
		panic("failed to connect database")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	// Pass the database connection to the handler
	loginHandler := handler.LoginHandler(db)
	// Define routes

	r.Post("/login", loginHandler)
	http.ListenAndServe(":8080", r)
}
