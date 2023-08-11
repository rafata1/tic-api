package model

import "time"

type Project struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ProjectMember struct {
	ID          int64     `db:"id"`
	ProjectID   int64     `db:"project_id"`
	MemberEmail string    `db:"member_email"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
