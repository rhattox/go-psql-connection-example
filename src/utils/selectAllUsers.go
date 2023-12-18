package utils

import (
	"database/sql"
	"fmt"
)
type User struct {
	ID       int
	Username string
	Email    string
}


func SelectAllUsers(db *sql.DB) ([]User, error) {
	// SQL statement to select all rows from 'app_user' table for PostgreSQL
	selectSQL := "SELECT id, username, email FROM app_user"

	// Execute the SQL statement
	rows, err := db.Query(selectSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result set and populate the slice of users
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	fmt.Println(users)
	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}