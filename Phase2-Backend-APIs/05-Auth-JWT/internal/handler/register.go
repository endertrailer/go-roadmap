package handler

import (
	"encoding/json"
	"net/http"

	"jwt_learn/internal/repository"
	"jwt_learn/internal/security"
)

type IncomingUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthHandler struct {
	Repo *repository.PostgresRepo
}

func (h *AuthHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req IncomingUser

	ctx := r.Context()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}
	if req.Username == "" || req.Password == "" {
		http.Error(w, "Empty username or password ", http.StatusBadRequest)
		return
	}
	hash, err := security.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "error hashing password", http.StatusBadRequest)
		return
	}

	if err := h.Repo.CreateUser(ctx, req.Username, hash); err != nil {
		http.Error(w, "Error Creating User", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

