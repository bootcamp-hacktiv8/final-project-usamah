package response

import (
	"final-project-usamah/entities"
	"time"
)

type LoginFormatter struct {
	Token string `json:"token"`
}

func ResponseLogin(token string) LoginFormatter {
	formatter := LoginFormatter{
		Token: token,
	}
	return formatter
}

type UserFormatter struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func ResponseUser(user entities.User) UserFormatter {
	formatter := UserFormatter{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}

	return formatter
}

type UpdateUserFormatter struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Age        int       `json:"age"`
	Updated_at time.Time `json:"updated_at"`
}

func ResponseUpdateUser(user entities.User) UpdateUserFormatter {
	formatter := UpdateUserFormatter{
		Id:         user.Id,
		Username:   user.Username,
		Email:      user.Email,
		Age:        user.Age,
		Updated_at: user.Updated_at,
	}
	return formatter
}
