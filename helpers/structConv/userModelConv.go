package structConv

import (
	db "WGameBack/database/db"
	"WGameBack/models"
)

func ToUserModel(user db.User) models.User {
	return models.User{
		UserID:       user.UserID,
		UserName:     user.UserName,
		Password:     user.Password,
		UserType:     user.UserType,
		Email:        user.Email,
		Token:        user.Token,
		RefreshToken: user.RefreshToken,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func ToCreateUserParams(user models.User) db.CreateUserParams {
	return db.CreateUserParams{
		UserID:       user.UserID,
		UserName:     user.UserName,
		Password:     user.Password,
		UserType:     user.UserType,
		Email:        user.Email,
		Token:        user.Token,
		RefreshToken: user.RefreshToken,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
