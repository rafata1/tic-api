package repository

import (
	"context"
	"github.com/rafata1/tic-api/model"
)

type IRepo interface {
	InsertProject(ctx context.Context, project model.Project) (int64, error)
	InsertProjectMember(ctx context.Context, projectMember model.ProjectMember) error
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

func NewProjectRepo() IRepo {
	return &repo{}
}
