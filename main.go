package main

import (
	"database_connector/database"
	"database_connector/handlers"
	"log"
	"net/http"
)

var db *database.Database

func init() {
	var err error
	db, err = database.NewDatabase("database.db")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var err error

	err = db.CreateUserTable()
	if err != nil {
		log.Fatal(err)
	}

	err = db.CreateModelTable()
	if err != nil {
		log.Fatal(err)
	}

	err = db.CreateTaskTable()
	if err != nil {
		log.Fatal(err)
	}

	// Creaete
	http.HandleFunc("/createUser", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUserHandler(w, *r, db)
	})

	http.HandleFunc("/createModel", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateModelHandler(w, *r, db)
	})

	http.HandleFunc("/createTable", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTaskHandler(w, *r, db)
	})

	// get
	http.HandleFunc("/models", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetModelHandler(w, *r, db)
	})

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTaskHandler(w, *r, db)
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUserHandler(w, *r, db)
	})

	log.Printf("Server is starting and listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	defer db.Close()
}
