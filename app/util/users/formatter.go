package users

import (
	userTypeEntity "service-api/app/core/user_type/entities"
	"service-api/app/core/users/entities"
)

type UserFormatter struct {
	ID       int                     `json:"id"`
	Name     string                  `json:"name"`
	Username string                  `json:"username"`
	Email    string                  `json:"email"`
	UserType userTypeEntity.UserType `json:"user_type"`
	Password string                  `json:"password"`
	Token    string                  `json:"token"`
}

func FormatUser(user entities.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       int(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Token:    token,
		Password: user.Password,
		Username: user.Username,
		UserType: user.UserType,
	}

	return formatter
}

func FormatUserRegister(user entities.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       int(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Token:    token,
		Username: user.Username,
		UserType: user.UserType,
	}

	return formatter
}
