package service

import (
	"context"
	"project-ta/entity"
	"project-ta/repository"

	"gorm.io/gorm"
)

type LayananServiceInj interface {
	AddLayanan(ctx context.Context, layananRequest entity.LayananRequest) (entity.Layanan, error)
	FindLayananById(ctx context.Context, id int) (entity.Layanan, error)
	DeleteLayananById(ctx context.Context, id int) error
	FindAllLayanan(ctx context.Context) ([]entity.Layanan, error)
	EditLayananById(ctx context.Context, id int, layananReq entity.LayananRequest) (entity.Layanan, error)
}

type LayananService struct {
	DB          *gorm.DB
	LayananRepo repository.LayananRepositoryInj
}

func NewLayananService(db *gorm.DB, lr repository.LayananRepositoryInj) LayananServiceInj {
	return LayananService{DB: db, LayananRepo: lr}
}

func (ls LayananService) AddLayanan(ctx context.Context, layananRequest entity.LayananRequest) (entity.Layanan, error) {
	tx := ls.DB.Begin()

	newLayanan, err := ls.LayananRepo.AddLayanan(ctx, layananRequest, tx)
	if err != nil {
		tx.Rollback()
		return entity.Layanan{}, err
	}

	tx.Commit()
	return newLayanan, nil
}

func (ls LayananService) FindLayananById(ctx context.Context, id int) (entity.Layanan, error) {
	tx := ls.DB.Begin()

	layanan, err := ls.LayananRepo.FindById(ctx, id, tx)
	if err != nil {
		tx.Rollback()
		return entity.Layanan{}, err
	}

	tx.Commit()
	return layanan, nil
}

func (ls LayananService) EditLayananById(ctx context.Context, id int, layananReq entity.LayananRequest) (entity.Layanan, error) {
	tx := ls.DB.Begin()

	editedLayanan, err := ls.LayananRepo.EditById(ctx, id, tx, layananReq)
	if err != nil {
		tx.Rollback()
		return entity.Layanan{}, err
	}

	tx.Commit()
	return editedLayanan, nil
}

func (ls LayananService) DeleteLayananById(ctx context.Context, id int) error {
	tx := ls.DB.Begin()

	err := ls.LayananRepo.DeleteById(ctx, id, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (ls LayananService) FindAllLayanan(ctx context.Context) ([]entity.Layanan, error) {
	tx := ls.DB.Begin()

	allLayanan, err := ls.LayananRepo.FindAll(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return allLayanan, nil
}
