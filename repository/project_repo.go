package repository

import (
	"context"
	"github.com/rafata1/tic-api/model"
)

type IProjectRepo interface {
	InsertProject(ctx context.Context, project model.Project) (int64, error)
	InsertProjectMember(ctx context.Context, projectMember model.ProjectMember) error
	GetProjectByID(ctx context.Context, id int64) (model.Project, error)
	GetProjectsByUser(ctx context.Context, email string) ([]model.Project, error)
	GetProjectMember(ctx context.Context, email string, projectID int64) (model.ProjectMember, error)
	AddFAQ(ctx context.Context, faq model.FAQ) (int64, error)
	ListFAQs(ctx context.Context, projectID int64) ([]model.FAQ, error)
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

var getProjectByID = `
	SELECT id, name, created_at, updated_at FROM project
	WHERE id = ?
`

func (r repo) GetProjectByID(ctx context.Context, id int64) (model.Project, error) {
	var res model.Project
	err := GetReadonly(ctx).GetContext(ctx, &res, getProjectByID, id)
	return res, err
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

var addFAQ = `
	INSERT INTO faq (project_id, question, answer)
	VALUES (:project_id, :question, :answer)
`

func (r repo) AddFAQ(ctx context.Context, faq model.FAQ) (int64, error) {
	res, err := GetTx(ctx).NamedExecContext(ctx, addFAQ, faq)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

var getProjectMember = `
	SELECT id, project_id, member_email FROM project_member
	WHERE member_email = ? AND project_id = ?
`

func (r repo) GetProjectMember(ctx context.Context, email string, projectID int64) (model.ProjectMember, error) {
	var res model.ProjectMember
	err := GetReadonly(ctx).GetContext(ctx, &res, getProjectMember, email, projectID)
	return res, err
}

var listFAQs = `
	SELECT id, question, answer, created_at, updated_at FROM faq
	WHERE project_id = ?
`

func (r repo) ListFAQs(ctx context.Context, projectID int64) ([]model.FAQ, error) {
	var res []model.FAQ
	err := GetReadonly(ctx).SelectContext(ctx, &res, listFAQs, projectID)
	return res, err
}

func NewProjectRepo() IProjectRepo {
	return &repo{}
}
