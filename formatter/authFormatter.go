package formatter

import (
	"myfinnplan/entity"
)

type UserFormatter struct {
	User  entity.User `json:"profile"`
	Token string      `json:"token"`
}

func FormatUser(user entity.User, token string) UserFormatter {
	formatter := UserFormatter{
		User:  user,
		Token: token,
	}
	formatter.User.Password = ""
	return formatter
}
