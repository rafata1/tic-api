package project

import "time"

type DTOProject struct {
	Name string `json:"name"`
}

type OutputProject struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
