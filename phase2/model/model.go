package model

type Todo struct {
	ID     int     `json:""`
	Title  *string `json:"title"`
	Status *bool   `json:"status"`
}
