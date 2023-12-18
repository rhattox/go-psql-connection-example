package utils

import "database/sql"

func CreateTableIfNotExists(db *sql.DB) error {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS app_user (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL,
			email VARCHAR(50) NOT NULL
		);
	`

	// Execute the SQL statement
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}
