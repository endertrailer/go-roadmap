package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello weebs")
}

func playerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello player")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /player", playerHandler)
	http.ListenAndServe(":8080", mux)
}
