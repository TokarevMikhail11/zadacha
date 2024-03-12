package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Expression struct {
	ID         int
	Expression string
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func addExpressionHandler(w http.ResponseWriter, r *http.Request) {
	var exp Expression
	err := json.NewDecoder(r.Body).Decode(&exp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save expression to database and return ID
	// Update status to "accepted" and store in DB
	exp.Status = "accepted"
	// Add expression to DB
}

func getExpressionsHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve all expressions from DB and return as JSON
}

func getExpressionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Retrieve expression by ID from DB and return as JSON
}

func getOperationsHandler(w http.ResponseWriter, r *http.Request) {
	// Return list of operations with their execution times as JSON
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Assign a task to the agent for execution and return as JSON
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	var result map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/add-expression", addExpressionHandler).Methods("POST")
	r.HandleFunc("/get-expressions", getExpressionsHandler).Methods("GET")
	r.HandleFunc("/get-expression/{id}", getExpressionHandler).Methods("GET")
	r.HandleFunc("/get-operations", getOperationsHandler).Methods("GET")
	r.HandleFunc("/get-task", getTaskHandler).Methods("GET")
	r.HandleFunc("/result", resultHandler).Methods("POST")

	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
