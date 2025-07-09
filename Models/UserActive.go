package Models

import "time"

type Post struct {
	Username  string    `json:"username" validate:"required"`
	Title     string    `json:"title" validate:"required,min=3,max=100"`
	Content   string    `json:"content" validate:"required,min=10"`
	CreatedAt time.Time `json:"created_at"`
}

type Like struct {
	PostID int `json:"PostID"`
	UserID int `json:"UserID"`
}
