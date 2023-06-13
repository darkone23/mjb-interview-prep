package user

import (
	"mjb-interview-prep/models"
)

type UserDto struct {
	Username string `json:"username" xml:"username" binding:"required"`
	// server generated argon2 encoded password hash
	// why argon2 encoded? because we will use the encoded salt to verify
	// against the user supplied key
	Password string `json:"password" xml:"password" binding:"required"`
}

func (dto UserDto) ToModel() *models.User {
	return &models.User{
		Username: dto.Username,
		Password: dto.Password,
	}
}

func FromModel(model *models.User) UserDto {
	return UserDto{
		Username: model.Username,
		Password: model.Password,
	}
}
