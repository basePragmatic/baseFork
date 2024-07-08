package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Job struct {
	ID          int
	Title       string
	Description string
	Salary      int
}

func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	createTable(db)
	return db, nil
}

func createTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS jobs (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "title" TEXT,
        "description" TEXT,
        "salary" INTEGER
    );`
	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
}

func InsertJob(db *sql.DB, title, description string, salary int) error {
	insertJobSQL := `INSERT INTO jobs (title, description, salary) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertJobSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(title, description, salary)
	return err
}
