package db

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)
 
var DB *sql.DB

 
func InitDB() {
    var err error
    DB, err = sql.Open("sqlite", "api.db")
    if err != nil {
        panic("Could not connect to database.")
    }
 
    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5)
 
    createTable()
}
 
func createTable() {
    createUsers := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    )
    `
    _, err := DB.Exec(createUsers)
    if err != nil {
        panic("Could not create users table.")
    }

    createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime DATETIME NOT NULL,
        user_id INTEGER, 
        FOREIGN KEY(user_id) REFERENCES users(id)
    )
    `
 
    _, err = DB.Exec(createEventsTable)
 
    if err != nil {
        panic("Could not create events table.")
    }

    RegisterTable := `
    CREATE TABLE IF NOT EXISTS registration(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        event_id INTEGER,
        user_id INTEGER,
        FOREIGN KEY(event_id) REFERENCES events(id),
        FOREIGN KEY(user_id) REFERENCES users(id)
    )
    `
    _, err = DB.Exec(RegisterTable)
    if err != nil {
        panic("Could not register user")
    }
}