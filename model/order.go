package model

type Order struct {
	Id           int    `json:"id"`
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	ConfirmEmail string `json:"confirm_email" validate:"required,eqfield=Email"`
	Message      string `json:"message"`
}
