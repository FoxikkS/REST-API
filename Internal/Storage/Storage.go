package Storage

import (
	"REST-API-pet-proj/Internal/Storage/Sqlite"
	"REST-API-pet-proj/Structure"
)

var Storage Sqlite.Storage

func SaveUser(username, email, passwordHash string) error {
	_, err := Storage.DB.Exec(
		`INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)`,
		username, email, passwordHash)
	return err
}

func GetUserPassword(username, email string) (string, error) {
	var passwordHash string
	err := Storage.DB.QueryRow(`
	SELECT password_hash 
	FROM users 
	WHERE email = ? AND username = ?`,
		email, username).Scan(&passwordHash)
	
	if err != nil {
		return "", err
	}
	return passwordHash, nil
}

func GetUserData(username string) (*Structure.UserData, error) {
	var userData Structure.UserData
	err := Storage.DB.QueryRow(`
	SELECT id, username, email, avatar_url, created_at 
	FROM users 
	WHERE username = ?`, username).Scan(
		&userData.ID, &userData.Username, &userData.Email,
		&userData.AvatarURL, &userData.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &userData, nil
}
