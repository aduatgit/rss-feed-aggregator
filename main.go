package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/aduatgit/rss-feed-aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL environment variable is not set")
	}
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	apiCfg := apiConfig{
		DB: database.New(db),
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

	rV1.Post("/users", apiCfg.handlerCreateUser)
	rV1.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	rV1.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	rV1.Get("/feeds", apiCfg.handlerGetFeeds)

	rV1.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollow))
	rV1.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	rV1.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Printf("Serving on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
