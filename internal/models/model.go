package models

import (
	"database/sql"
	"fmt"
)

// Model represents a data model in the application.
type Model struct {
	ID   int
	Name string
}

// GetAllModels retrieves all models from the database.
func GetAllModels(db *sql.DB) ([]Model, error) {
	query := "SELECT id, name FROM models"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var models []Model
	for rows.Next() {
		var model Model
		err := rows.Scan(&model.ID, &model.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		models = append(models, model)
	}

	return models, nil
}

// CreateModel creates a new model in the database.
func CreateModel(db *sql.DB, model Model) error {
	query := "INSERT INTO models (name) VALUES ($1)"
	_, err := db.Exec(query, model.Name)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

// UpdateModel updates an existing model in the database.
func UpdateModel(db *sql.DB, model Model) error {
	query := "UPDATE models SET name = $1 WHERE id = $2"
	_, err := db.Exec(query, model.Name, model.ID)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

// DeleteModel deletes a model from the database.
func DeleteModel(db *sql.DB, id int) error {
	query := "DELETE FROM models WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}