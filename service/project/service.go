package project

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/rafata1/tic-api/model"
	"github.com/rafata1/tic-api/repository"
	"github.com/rafata1/tic-api/service/common"
	"log"
)

type IService interface {
	CreateProject(ctx context.Context, name string) (OutputProject, error)
}

type service struct {
	txnProvider repository.Provider
	projectRepo repository.IRepo
}

func (s service) CreateProject(ctx context.Context, name string) (OutputProject, error) {
	if name == "" {
		return OutputProject{}, ErrProjectNameRequired
	}
	var projectID int64
	var err error
	err = s.txnProvider.Transact(ctx, func(ctx context.Context) error {
		projectID, err = s.projectRepo.InsertProject(
			ctx,
			model.Project{Name: name},
		)
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

func NewService(db *sqlx.DB) IService {
	return &service{
		txnProvider: repository.NewProvider(db),
		projectRepo: repository.NewProjectRepo(),
	}
}
