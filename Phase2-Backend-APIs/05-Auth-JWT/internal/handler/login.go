package handler

import (
	"encoding/json"
	"net/http"

	"jwt_learn/internal/auth"
	"jwt_learn/internal/security"
)

type LoginRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type LoginResponse struct {
	ResponseToken string `json:"responseToken"`
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	loginRequestStruct := LoginRequest{}

	err := json.NewDecoder(r.Body).Decode(&loginRequestStruct)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	databaseUser, err := h.Repo.GetUserByUsername(r.Context(), loginRequestStruct.Username)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	validity := security.VerifyPassword(loginRequestStruct.Password, databaseUser.Hashedpassword)
	// TODO: What if 'validity' is false (wrong password)? Currently, if validity is false, it skips this if-block and ends without sending an HTTP error response! Add an 'else' block (or check 'if !validity { ... return }') to return http.StatusUnauthorized.
	if validity {
		response, err := auth.GenerateToken(databaseUser.ID)
		responseToken := LoginResponse{ResponseToken: response}

		if err != nil {
			http.Error(w, "Error generating Token please try again", http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(responseToken); err != nil {
			// TODO: Haha, love the error message! Note that since w.WriteHeader was already called above, sending an http.Error here won't work well. Logging the error with log.Printf is usually best here!
			http.Error(w, "Idk what to call this error", http.StatusBadRequest)
		}
	}
}
