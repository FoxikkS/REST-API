package Sqlite

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

type Storage struct {
	DB *sql.DB
}

func InitStorage(storagePath string) (*Storage, error) {
	// пока заюзаю склайт, потом мигрирую на постгрес
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
	);

	CREATE TABLE IF NOT EXISTS posts (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	user_id INTEGER NOT NULL,
    	title TEXT NOT NULL,
    	content TEXT NOT NULL,
    	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    	FOREIGN KEY(user_id) REFERENCES users(id)
	);

	CREATE TABLE IF NOT EXISTS likes (
    	user_id INTEGER NOT NULL,
    	post_id INTEGER NOT NULL,
    	FOREIGN KEY(user_id) REFERENCES users(id),
    	FOREIGN KEY(post_id) REFERENCES posts(id),
    	UNIQUE(user_id, post_id)
	);
`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, err
	}

	return &Storage{DB: db}, nil
}
