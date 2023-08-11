package model

import "time"

type FAQ struct {
	ID        int64     `db:"id"`
	ProjectID int64     `db:"project_id"`
	Question  string    `db:"question"`
	Answer    string    `db:"answer"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
