package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"github.com/rhattox/go-psql-connection-example/utils"

)

func main() {

	db, err := sql.Open("postgres", "postgres://postgres:postgres@db:5432/psql_connection_example?sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Check if the 'user' table exists
	if err := utils.CreateTableIfNotExists(db); err != nil {
		log.Fatal(err)
	}

	// Manually specified users
	users := []struct {
		Username string
		Email    string
	}{
		{"user1", "user1@example.com"},
		{"user2", "user2@example.com"},
		// Add more users as needed
	}

	// Insert users into the 'user' table
	for _, user := range users {
		if err := utils.InsertUser(db, user.Username, user.Email); err != nil {
			log.Fatal(err)
		}
	}

	utils.SelectAllUsers(db)

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
