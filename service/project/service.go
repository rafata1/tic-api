package project

import (
	"context"
	"git.teko.vn/dung.cda/tic-26-be/model"
	"git.teko.vn/dung.cda/tic-26-be/repository"
	"git.teko.vn/dung.cda/tic-26-be/service/common"
	"github.com/jmoiron/sqlx"
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
