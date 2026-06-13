package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/tomiHapvaDosta/ZedFirstTry/internal/database"
)

type apiConfig struct {
	tokenSecret string
	queries     *database.Queries
}

func main() {
	godotenv.Load()

	tokenSecret := os.Getenv("SECRET")
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Print(err.Error())
		return
	}

	dbQueries := database.New(db)

	const filePathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()

	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	cfg := apiConfig{tokenSecret: tokenSecret, queries: dbQueries}

	mux.HandleFunc("POST /api/users", cfg.createUser)
	mux.HandleFunc("POST /api/login", cfg.loginUser)
	mux.HandleFunc("POST /api/posts", cfg.createPost)
	mux.HandleFunc("GET /api/posts", cfg.getPosts)
	mux.HandleFunc("GET /api/posts/{id}", cfg.getPost)
	mux.HandleFunc("PUT /api/posts/{id}", cfg.changePost)
	mux.HandleFunc("DELETE /api/posts/{id}", cfg.deletePost)
	mux.HandleFunc("POST /api/posts/{id}/publish", cfg.publishPost)

	log.Fatal(server.ListenAndServe())
}
