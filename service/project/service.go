package project

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/rafata1/tic-api/model"
	"github.com/rafata1/tic-api/repository"
	"github.com/rafata1/tic-api/service/auth"
	"github.com/rafata1/tic-api/service/common"
	"log"
)

type IService interface {
	CreateProject(ctx context.Context, name string) (outputProject, error)
	ListProject(ctx context.Context) ([]outputProject, error)
	AddFAQ(ctx context.Context, input addFAQInput) (outputFAQ, error)
}

type service struct {
	txnProvider repository.Provider
	projectRepo repository.IProjectRepo
}

func (s service) CreateProject(ctx context.Context, name string) (outputProject, error) {
	if name == "" {
		return outputProject{}, ErrProjectNameRequired
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
		return outputProject{}, common.ErrExecuteIntoDB
	}

	return outputProject{
		ID:   projectID,
		Name: name,
	}, nil
}

func (s service) ListProject(ctx context.Context) ([]outputProject, error) {
	userEmail := auth.GetUserEmail(ctx)
	ctx = s.txnProvider.Readonly(ctx)
	projects, err := s.projectRepo.GetProjectsByUser(ctx, userEmail)
	if err != nil {
		log.Printf("error listing projects %s\n", err.Error())
		return nil, common.ErrQueryIntoDB
	}
	return toOutputProjects(projects), err
}

func toOutputProjects(projects []model.Project) []outputProject {
	res := make([]outputProject, 0, len(projects))
	for _, project := range projects {
		res = append(res, toOutputProject(project))
	}
	return res
}

func toOutputProject(project model.Project) outputProject {
	return outputProject{
		ID:        project.ID,
		Name:      project.Name,
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
	}
}

func (s service) AddFAQ(ctx context.Context, input addFAQInput) (outputFAQ, error) {
	var err error
	var id int64
	err = s.txnProvider.Transact(ctx, func(ctx context.Context) error {
		err = s.validateProjectPerm(ctx, input.ProjectID)
		if err != nil {
			return err
		}
		id, err = s.projectRepo.AddFAQ(ctx, toFAQModel(input))
		return err
	})
	return outputFAQ{ID: id}, err
}

func toFAQModel(input addFAQInput) model.FAQ {
	return model.FAQ{
		ProjectID: input.ProjectID,
		Question:  input.Question,
		Answer:    input.Answer,
	}
}

func (s service) validateProjectPerm(ctx context.Context, projectID int64) error {
	userEmail := auth.GetUserEmail(ctx)
	projectMember, err := s.projectRepo.GetProjectMember(ctx, userEmail, projectID)
	if err == sql.ErrNoRows {
		return common.ErrUnauthorized
	}

	if err != nil {
		log.Printf("error get project member %s\n", err.Error())
		return common.ErrQueryIntoDB
	}

	if projectMember.ID <= 0 {
		return common.ErrUnauthorized
	}
	return nil
}

func NewService(db *sqlx.DB) IService {
	return &service{
		txnProvider: repository.NewProvider(db),
		projectRepo: repository.NewProjectRepo(),
	}
}
