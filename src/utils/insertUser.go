package utils

import "database/sql"

func InsertUser(db *sql.DB, username, email string) error {
	// SQL statement to insert into 'user' table
	insertSQL := "INSERT INTO app_user (username, email) VALUES ($1, $2)"

	// Execute the SQL statement
	_, err := db.Exec(insertSQL, username, email)
	if err != nil {
		return err
	}

	return nil
}

