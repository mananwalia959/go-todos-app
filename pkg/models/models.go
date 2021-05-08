package models

import "github.com/google/uuid"

type Todo struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
}

type Todos []Todo

type ErrorResponse struct {
	Message string `json:"message"`
}
