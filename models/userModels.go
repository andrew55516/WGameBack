package models

import "time"

type User struct {
	//ID           int64     `json:"id"`
	UserName string `json:"user_name" validate:"required, min=1, max=100"`
	//FirstName    string    `json:"first_name" validate:"required, min=1, max=100"`
	//LastName     string    `json:"last_name" validate:"required, min=1, max=100"`
	Password string `json:"password" validate:"required, min=1, max=100"`
	Email    string `json:"email" validate:"email, required"`
	//Phone        string    `json:"phone" validate:"required"`
	Token        string    `json:"token"`
	UserType     string    `json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	UserID       string    `json:"user_id"`
}
