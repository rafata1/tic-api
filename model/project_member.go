package model

import "time"

type ProjectMember struct {
	ID          int       `db:"id"`
	ProjectID   int       `db:"project_id"`
	MemberEmail string    `db:"member_email"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"created_at"`
}
