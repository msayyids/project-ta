package service

import (
	"context"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/repository"

	"github.com/jmoiron/sqlx"
)

type LayananServiceInj interface {
	AddLayanan(ctx context.Context, layananRequest entity.LayananRequest) (entity.Layanan, error)
	FindLayananById(ctx context.Context, id int) (entity.Layanan, error)
	DeleteLayananById(ctx context.Context, id int) error
	FindAllLayanan(ctx context.Context) ([]entity.Layanan, error)
	EditLayananById(ctx context.Context, id int, layananReq entity.LayananRequest) (entity.Layanan, error)
}

type LayananService struct {
	DB          *sqlx.DB
	LayananRepo repository.LayananRepositoryInj
}

func NewLayananService(db *sqlx.DB, lr repository.LayananRepositoryInj) LayananServiceInj {
	return LayananService{DB: db, LayananRepo: lr}
}

func (ls LayananService) AddLayanan(ctx context.Context, layananRequest entity.LayananRequest) (entity.Layanan, error) {
	tx, err := ls.DB.Beginx()

	if err != nil {
		return entity.Layanan{}, err
	}
	defer helper.CommitOrRollback(tx)

	newLayanan, err := ls.LayananRepo.AddLayanan(ctx, layananRequest, *tx)
	if err != nil {
		return entity.Layanan{}, err
	}

	return newLayanan, nil

}

func (ls LayananService) FindLayananById(ctx context.Context, id int) (entity.Layanan, error) {
	tx, err := ls.DB.Beginx()

	if err != nil {
		return entity.Layanan{}, err
	}
	defer helper.CommitOrRollback(tx)

	layanan, err := ls.LayananRepo.FindById(ctx, id, *tx)
	if err != nil {
		return entity.Layanan{}, err
	}

	return layanan, nil
}

func (ls LayananService) EditLayananById(ctx context.Context, id int, layananReq entity.LayananRequest) (entity.Layanan, error) {
	tx, err := ls.DB.Beginx()
	if err != nil {
		return entity.Layanan{}, err
	}
	defer helper.CommitOrRollback(tx)

	editedLayanan, err := ls.LayananRepo.EditById(ctx, id, tx, layananReq)
	if err != nil {
		return entity.Layanan{}, err
	}

	return editedLayanan, nil
}

func (ls LayananService) DeleteLayananById(ctx context.Context, id int) error {
	tx, err := ls.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = ls.LayananRepo.DeleteById(ctx, id, tx)
	helper.PanicIfError(err)

	return nil

}

func (ls LayananService) FindAllLayanan(ctx context.Context) ([]entity.Layanan, error) {
	tx, err := ls.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	allLayanan, err := ls.LayananRepo.FindAll(ctx, *tx)
	helper.PanicIfError(err)

	return allLayanan, nil

}
