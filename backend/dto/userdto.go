package dto

import "Beeper/backend/model"

type Userdto struct {
	ID uint			`json:"id"`
	Username string	`json:"username"`
}

func ToUserdto(user model.User) Userdto {
	return Userdto{
		ID:user.ID,
		Username:user.Username,
	}
}