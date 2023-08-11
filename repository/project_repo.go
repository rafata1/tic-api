package repository

import (
	"context"
	"github.com/rafata1/tic-api/model"
)

type IProjectRepo interface {
	InsertProject(ctx context.Context, project model.Project) (int64, error)
	InsertProjectMember(ctx context.Context, projectMember model.ProjectMember) error
	GetProjectsByUser(ctx context.Context, email string) ([]model.Project, error)
}

type repo struct{}

var insertProjectMember = `
	INSERT INTO project_member (project_id, member_email)
	VALUES (:project_id, :member_email)
`

func (r repo) InsertProjectMember(ctx context.Context, projectMember model.ProjectMember) error {
	_, err := GetTx(ctx).NamedExecContext(ctx, insertProjectMember, projectMember)
	return err
}

var insertProject = `
	INSERT INTO project (name) VALUES (:name)
`

func (r repo) InsertProject(ctx context.Context, project model.Project) (int64, error) {
	res, err := GetTx(ctx).NamedExecContext(ctx, insertProject, project)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

var getProjectsByUser = `
	SELECT p.id, p.name, p.created_at, p.updated_at FROM project AS p
	JOIN project_member AS pm ON p.id = pm.project_id
	WHERE pm.member_email = ?
`

func (r repo) GetProjectsByUser(ctx context.Context, email string) ([]model.Project, error) {
	var res []model.Project
	err := GetReadonly(ctx).SelectContext(ctx, &res, getProjectsByUser, email)
	return res, err
}

func NewProjectRepo() IProjectRepo {
	return &repo{}
}
