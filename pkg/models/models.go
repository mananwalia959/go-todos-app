package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedOn   time.Time `json:"-"`
}

type TodoCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TodoEditRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type Todos []Todo

type ErrorResponse struct {
	Message string `json:"message"`
}
