package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/new", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("new one"))
	})
	r.Get("/add", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("add one"))
	})
	http.ListenAndServe(":8080", r)
}
