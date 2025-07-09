package Storage

import (
	"REST-API-pet-proj/Internal/Storage/Sqlite"
	"REST-API-pet-proj/Models"
	"fmt"
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

func GetUserData(username string) (*Models.UserData, error) {
	var userData Models.UserData
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

func CreatePost(username, title, content string) error {
	var userID int
	err := Storage.DB.QueryRow(`SELECT id FROM users WHERE username = ?`, username).Scan(&userID)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	_, err = Storage.DB.Exec(
		`INSERT INTO posts (user_id, title, content) VALUES (?, ?, ?)`,
		userID, title, content)
	return err
}

func PutALike(userId, postId int) error {
	_, err := Storage.DB.Exec(
		`INSERT INTO likes (user_id, post_id) VALUES (?, ?)`,
		userId, postId)
	if err != nil {
		return err
	}
	return nil
}
