package Models

type UserRegistration struct {
	Username string `json:"username" validate:"min=3,max=32"`
	Password string `json:"password" validate:"min=4,max=31"`
	Email    string `json:"email" validate:"email"`
}

type UserData struct {
	ID        int
	Username  string
	Email     string
	AvatarURL *string
	CreatedAt string
}
