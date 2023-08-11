package project

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/rafata1/tic-api/model"
	"github.com/rafata1/tic-api/repository"
	"github.com/rafata1/tic-api/service/auth"
	"github.com/rafata1/tic-api/service/common"
	"log"
)

type IService interface {
	CreateProject(ctx context.Context, name string) (OutputProject, error)
	ListProject(ctx context.Context) ([]OutputProject, error)
}

type service struct {
	txnProvider repository.Provider
	projectRepo repository.IProjectRepo
}

func (s service) CreateProject(ctx context.Context, name string) (OutputProject, error) {
	if name == "" {
		return OutputProject{}, ErrProjectNameRequired
	}
	var projectID int64
	var err error

	userEmail := auth.GetUserEmail(ctx)
	err = s.txnProvider.Transact(ctx, func(ctx context.Context) error {
		projectID, err = s.projectRepo.InsertProject(
			ctx,
			model.Project{Name: name},
		)
		if err != nil {
			return err
		}

		err = s.projectRepo.InsertProjectMember(ctx, model.ProjectMember{
			ProjectID:   projectID,
			MemberEmail: userEmail,
		})
		return err
	})

	if err != nil {
		log.Printf("errors creating project: %s\n", err.Error())
		return OutputProject{}, common.ErrExecuteIntoDB
	}

	return OutputProject{
		ID:   projectID,
		Name: name,
	}, nil
}

func (s service) ListProject(ctx context.Context) ([]OutputProject, error) {
	userEmail := auth.GetUserEmail(ctx)
	ctx = s.txnProvider.Readonly(ctx)
	projects, err := s.projectRepo.GetProjectsByUser(ctx, userEmail)
	if err != nil {
		log.Printf("error listing projects %s\n", err.Error())
		return nil, common.ErrQueryIntoDB
	}
	return toOutputProjects(projects), err
}

func toOutputProjects(projects []model.Project) []OutputProject {
	res := make([]OutputProject, 0, len(projects))
	for _, project := range projects {
		res = append(res, toOutputProject(project))
	}
	return res
}

func toOutputProject(project model.Project) OutputProject {
	return OutputProject{
		ID:        project.ID,
		Name:      project.Name,
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
	}
}

func NewService(db *sqlx.DB) IService {
	return &service{
		txnProvider: repository.NewProvider(db),
		projectRepo: repository.NewProjectRepo(),
	}
}
