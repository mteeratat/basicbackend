package model

type Todo struct {
	ID     int
	Title  *string `json:"title" validate:"required"`
	Status *bool   `json:"status" validate:"required"`
}
