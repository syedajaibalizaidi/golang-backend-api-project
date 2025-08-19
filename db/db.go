// to establish the connection bw tables we foriegn key

package db

import (
	"database/sql" // but we will interact with built in sql pkg which is part of go standard library

	_ "github.com/mattn/go-sqlite3" // go uses it under the hood we wont use it directly , TO MAKE SURE THIS FILE WONT GET DISAPPEAR WHEN WE SAVE IT SO WE ADD _ BEFORE IT. WHICH TELLS WE NEED THIS IMPORT.
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db") // to open a connection which wants driver name which sqlite3 in case of datasource name we need to pass the path

	if err != nil {
		panic("database connection failed.") // when we dont have any connection we will crash the app.
	}

	DB.SetMaxOpenConns(10) // sets how many open connections we can have
	DB.SetMaxIdleConns(5)  // how much connections ca n remain idle. if no one using these connections.

	createTables()
}

// creating db table
func createTables() {

	createUsersTables := `
	CREATE TABLES IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL, 
	)
	`

	_, err := DB.Exec(createUsersTables)

	if err != nil {
		panic("user creation failed")
	}

	createEventsTables := `
	CREATE TABLES IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT, 
	name TEXT NOT NULL, 
	description TEXT NOT NULL, 
	loaction TEXT NOT NULL, 
	dateTime DATETIME, 
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id) 
	)
	`

	_, err = DB.Exec(createEventsTables)

	if err != nil {
		panic("cant create the events table")
	}

	createRegistrationTable := `
    CREATE TABLE IF NOT EXIST registrations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_id INTEGER,
    user_id INTEGER,
    FOREIGN KEY(event_id) REFERENCES events(id),
    FOREIGN KEY(user_id) REFERENCES users(id),
    )
    `

	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		panic("cant create the registration table")
	}
}
