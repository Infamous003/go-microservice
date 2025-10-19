package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger) // Adding a logger middleware

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("What's up!"))
	})

	s := &http.Server{
		Addr:         ":9090",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  30 * time.Minute,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server listening http://localhost:9090")
	s.ListenAndServe()
}
