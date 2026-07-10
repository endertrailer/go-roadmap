package main

import (
	"encoding/json"
	"net/http"
)

type player struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

func playerHandler(w http.ResponseWriter, r *http.Request) {
	p1 := player{ID: 1, Name: "mysteries", Level: 70}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /player", playerHandler)
	http.ListenAndServe(":8080", mux)
}
