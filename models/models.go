package models

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Note        string `json:"note"`
}
