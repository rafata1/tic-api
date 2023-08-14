package project

import "time"

type inputProject struct {
	Name string `json:"name"`
}

type outputProject struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type inputFAQ struct {
	ProjectID int64  `json:"project_id"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
}

type outputFAQ struct {
	ID        int64     `json:"id"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
