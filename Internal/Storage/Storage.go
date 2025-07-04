package Storage

import (
	"REST-API-pet-proj/Internal/Storage/Sqlite"
)

var Storage Sqlite.Storage

func SaveUser(username, email, password_hash string) error {
	_, err := Storage.DB.Exec(
		`INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)`,
		username, email, password_hash)
	return err
}
