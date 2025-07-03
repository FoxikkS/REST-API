package Sqlite

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

type Storage struct {
	DB *sql.DB
}

func InitStorage(storagePath string) (*Storage, error) {
	db, err := sql.Open("sqlite", storagePath)
	if err != nil {
		return nil, err
	}
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS users (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	username TEXT NOT NULL UNIQUE,
    	email TEXT NOT NULL UNIQUE,
    	password_hash TEXT NOT NULL,
    	avatar_url TEXT,
    	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, err
	}

	return &Storage{DB: db}, nil
}
