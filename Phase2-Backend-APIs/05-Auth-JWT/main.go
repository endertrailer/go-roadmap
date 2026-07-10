package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"jwt_learn/internal/auth"
	"jwt_learn/internal/handler"
	"jwt_learn/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	ctx := context.Background()

	connStr := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Printf("warning: %v", err)
	}
	defer pool.Close()
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Database ping failed: %v\n", err)
	}
	repo := repository.NewPostgresRepo(pool)
	authHandler := &handler.AuthHandler{Repo: repo}
	mux := http.NewServeMux()
	mux.HandleFunc("POST /createUser", authHandler.CreateUserHandler)
	mux.HandleFunc("POST /login", authHandler.LoginHandler)

	// Here is how you wrap a protected route with AuthMiddleware!
	// Notice that AuthMiddleware takes the handler function as an argument:
	mux.HandleFunc("GET /profile", auth.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Protected route accessed successfully! Authenticated User ID: " + fmt.Sprint(userID)))
	}))

	port := os.Getenv("port")
	if port == "" {
		port = "8080" // Default fallback
	}
	http.ListenAndServe(":"+port, mux)
}
