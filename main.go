package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	rV1 := chi.NewRouter()
	r.Mount("/v1", rV1)

	rV1.Get("/healthz", handlerReadiness)
	rV1.Get("/err", handlerError)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Printf("Serving on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
