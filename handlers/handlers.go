package handlers

import (
	"database_connector/database"
	"database_connector/people"
	"encoding/json"
	"fmt"
	"io"
	_ "log"
	"net/http"
)

// Post handlers

func CreateUserHandler(w http.ResponseWriter, r http.Request, db *database.Database) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var user people.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	err = db.CreateUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Error creating user in database", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func CreateModelHandler(w http.ResponseWriter, r http.Request, db *database.Database) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var model people.Model
	err = json.Unmarshal(body, &model)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	err = db.CreateModel(model.ModelName, model.UserId)
	if err != nil {
		http.Error(w, "Error creating model in database", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func CreateTaskHandler(w http.ResponseWriter, r http.Request, db *database.Database) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var task people.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	err = db.CreateTask(task.Context, task.IsCompleted, task.ModelId)
	if err != nil {
		http.Error(w, "Error creating task in database", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Get handlers
func GetModelHandler(w http.ResponseWriter, r http.Request, db *database.Database) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	userID := r.URL.Query().Get("userid")
	if userID == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	models, err := db.GetAllModels(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(models)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}

func GetTaskHandler(w http.ResponseWriter, r http.Request, db *database.Database) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	modelID := r.URL.Query().Get("modelid")
	if modelID == "" {
		http.Error(w, "Model id is required", http.StatusBadRequest)
		return
	}

	tasks, err := db.GetAllTasks(modelID)
	if err != nil {
		http.Error(w, "Model not found", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	jsonData, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)

}

func GetUserHandler(w http.ResponseWriter, r http.Request, db *database.Database) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}
	user, err := db.GetUser(username)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
