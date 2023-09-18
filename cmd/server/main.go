package main

import (
	"log"
	"net/http"

	"github.com/Sotnasjeff/cep-multithreading-api/internal/webserver/handlers"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Use(LogRequest)
	r.Route("/cep", func(r chi.Router) {
		r.Get("/{cep}", handlers.GetAddress)
	})

	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n\nRequest enviada %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
