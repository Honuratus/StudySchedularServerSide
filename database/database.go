package database

import (
	"database/sql"
	"database_connector/people"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	conn *sql.DB
}

func NewDatabase(path string) (*Database, error) {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	db := &Database{conn: conn}
	return db, nil
}

func (db *Database) Close() error {
	return db.conn.Close()
}

func (db *Database) CreateUserTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE,
			password TEXT
		);
	`
	_, err := db.conn.Exec(query)
	return err
}

func (db *Database) CreateModelTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS models (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			model_name TEXT,
			user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`
	_, err := db.conn.Exec(query)
	return err
}

func (db *Database) CreateTaskTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			context TEXT,
			isCompleted INTEGER,
			model_id INTEGER,
			FOREIGN KEY (model_id) REFERENCES models(id)
		)
	`
	_, err := db.conn.Exec(query)
	return err
}

// Create Methods

func (db *Database) CreateUser(username string, password string) error {
	query := `
		INSERT INTO users (username, password)
		VALUES (?, ?)
	`
	_, err := db.conn.Exec(query, username, password)
	return err
}

func (db *Database) CreateModel(modelName string, userID int) error {
	query := `
		INSERT INTO models (model_name, user_id)
		VALUES (?, ?)
	`
	_, err := db.conn.Exec(query, modelName, userID)
	return err
}

func (db *Database) CreateTask(context string, isCompleted int, modelID int) error {
	query := `
		INSERT INTO tasks (context, is_completed, model_id)
		VALUES(?, ?, ?)
	`
	_, err := db.conn.Exec(query, context, isCompleted, modelID)
	return err
}

// Get Methods
func (db *Database) GetAllModels(userID string) ([]people.Model, error) {
	var models []people.Model
	query := `
		SELECT * FROM models WHERE user_id = ?
	`
	rows, err := db.conn.Query(query, userID)
	defer rows.Close()

	for rows.Next() {
		var model people.Model
		err := rows.Scan(&model.ID, &model.ModelName, &model.UserId)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return models, nil
}

func (db *Database) GetAllTasks(modelID string) ([]people.Task, error) {
	var tasks []people.Task
	query := `
		SELECT * FROM tasks WHERE model_id = ?
	`
	rows, err := db.conn.Query(query, modelID)

	for rows.Next() {
		var task people.Task
		err := rows.Scan(&task.ID, &task.Context, &task.IsCompleted, &task.ModelId)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (db *Database) GetUser(username string) (*people.User, error) {
	var user people.User
	query := `
		SELECT * FROM users WHERE username = ?
	`
	row := db.conn.QueryRow(query, username)

	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
