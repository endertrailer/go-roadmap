package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func handleIncoming(w http.ResponseWriter, r *http.Request) {
	p1 := player{ID: 1, Name: "klein", Level: 10}
	json.NewEncoder(w).Encode(p1)
}

func (a *app) insertData(w http.ResponseWriter, r *http.Request) {
	var dep dept
	err := json.NewDecoder(r.Body).Decode(&dep)
	if err != nil {
		http.Error(w, "malformed json", http.StatusBadRequest)
	}
	sql := "INSERT INTO dept (deptno, dname, loc) VALUES ($1, $2, $3)"
	_, err = a.db.Exec(context.Background(), sql, dep.Deptno, dep.Dname, dep.Loc)
}

func (a *app) updateData(w http.ResponseWriter, r *http.Request) {
	var dep dept
	err := json.NewDecoder(r.Body).Decode(&dep)
	if err != nil {
		http.Error(w, "malformed json", http.StatusBadRequest)
	}
	sql := "update dept set dname = $1, loc = $2 where deptno = $3"
	_, err = a.db.Exec(context.Background(), sql, dep.Dname, dep.Loc, dep.Deptno)
}

func (a *app) deleteRow(w http.ResponseWriter, r *http.Request) {
	var deptNo struct {
		Deptno int `json:"deptno"`
	}
	err := json.NewDecoder(r.Body).Decode(&deptNo)
	if err != nil {
		http.Error(w, "malformed json", http.StatusBadRequest)
	}
	sql := "delete from dept where deptno = $1"
	_, err = a.db.Exec(context.Background(), sql, deptNo.Deptno)
}

type app struct {
	db *pgxpool.Pool
}
type player struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}
type dept struct {
	Deptno int    `json:"deptno"`
	Dname  string `json:"dname"`
	Loc    string `json:"loc"`
}

func main() {
	// establish the connection to the database
	godotenv.Load()
	connStr := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(context.Background(), connStr)
	a := &app{db: pool}
	if err != nil {
		log.Fatal("failed to connect:", err)
	}
	rows, err := pool.Query(context.Background(), "select deptno,dname,loc from dept")
	if err != nil {
		log.Fatal("query failed:", err)
	}
	depts, err := pgx.CollectRows(rows, pgx.RowToStructByName[dept])
	for _, d := range depts {
		fmt.Printf("%d \n", d.Deptno)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handleIncoming)
	mux.HandleFunc("POST /dept", a.insertData)
	mux.HandleFunc("POST /deptUpdate", a.updateData)
	mux.HandleFunc("POST /deptDelete", a.deleteRow)
	http.ListenAndServe(":8080", mux)
}
