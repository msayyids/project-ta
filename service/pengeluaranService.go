package service

import (
	"context"
	"fmt"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/repository"

	"github.com/jmoiron/sqlx"
)

type PengeluaranService struct {
	DB              sqlx.DB
	PengeluaranRepo repository.PengeluaranRepositoryInj
}
type PengeluaranServiceInj interface {
	CreatePengeluaran(ctx context.Context, pengeluaranReq entity.PengeluaranRequest) (entity.Pengeluaran, error)
	EditPengeluaran(ctx context.Context, pengeluaranReq entity.PengeluaranRequest, id int) (entity.Pengeluaran, error)
	FindPengeluaranById(ctx context.Context, id int) (entity.Pengeluaran, error)
	FindAllPengeluaran(ctx context.Context) ([]entity.Pengeluaran, error)
	DeletePengeluaran(ctx context.Context, id int) error
}

func NewPengeluaranService(db sqlx.DB, pr repository.PengeluaranRepositoryInj) PengeluaranServiceInj {
	return PengeluaranService{
		DB:              db,
		PengeluaranRepo: pr,
	}
}

func (pr PengeluaranService) CreatePengeluaran(ctx context.Context, pengeluaranReq entity.PengeluaranRequest) (entity.Pengeluaran, error) {
	tx, err := pr.DB.Beginx()
	// helper.PanicIfError(err)
	if err != nil {
		return entity.Pengeluaran{}, fmt.Errorf("error disini")
	}

	defer helper.CommitOrRollback(tx)

	newPengeluaran, err := pr.PengeluaranRepo.AddPengeluaran(ctx, *tx, pengeluaranReq)
	// helper.PanicIfError(err)
	if err != nil {
		return entity.Pengeluaran{}, fmt.Errorf("error pengeluaran 2")
	}

	return newPengeluaran, nil

}

func (pr PengeluaranService) EditPengeluaran(ctx context.Context, pengeluaranReq entity.PengeluaranRequest, id int) (entity.Pengeluaran, error) {
	tx, err := pr.DB.Beginx()
	// helper.PanicIfError(err)
	if err != nil {
		return entity.Pengeluaran{}, fmt.Errorf("error disini")
	}

	defer helper.CommitOrRollback(tx)

	newPengeluaran, err := pr.PengeluaranRepo.UpdatePengeluaran(ctx, *tx, id, pengeluaranReq)
	// helper.PanicIfError(err)
	if err != nil {
		return entity.Pengeluaran{}, fmt.Errorf("error pengeluaran 2")
	}

	return newPengeluaran, nil

}

func (pr PengeluaranService) FindPengeluaranById(ctx context.Context, id int) (entity.Pengeluaran, error) {
	tx, err := pr.DB.Beginx()
	// helper.PanicIfError(err)
	if err != nil {
		return entity.Pengeluaran{}, fmt.Errorf("error disini")
	}

	defer helper.CommitOrRollback(tx)

	newPengeluaran, err := pr.PengeluaranRepo.FindPengeluaranById(ctx, *tx, id)
	// helper.PanicIfError(err)
	if err != nil {
		return entity.Pengeluaran{}, fmt.Errorf("error pengeluaran 2")
	}

	return newPengeluaran, nil

}

func (pr PengeluaranService) FindAllPengeluaran(ctx context.Context) ([]entity.Pengeluaran, error) {
	tx, err := pr.DB.Beginx()
	// helper.PanicIfError(err)
	if err != nil {
		return []entity.Pengeluaran{}, fmt.Errorf("error disini")
	}

	defer helper.CommitOrRollback(tx)

	newPengeluaran, err := pr.PengeluaranRepo.FindAllPengeluaran(ctx, *tx)
	// helper.PanicIfError(err)
	if err != nil {
		return []entity.Pengeluaran{}, fmt.Errorf("error pengeluaran 2")
	}

	return newPengeluaran, nil

}

func (pr PengeluaranService) DeletePengeluaran(ctx context.Context, id int) error {
	tx, err := pr.DB.Beginx()
	// helper.PanicIfError(err)
	if err != nil {
		return fmt.Errorf("error disini")
	}

	defer helper.CommitOrRollback(tx)

	err = pr.PengeluaranRepo.DeletePengeluaran(ctx, *tx, id)
	// helper.PanicIfError(err)
	if err != nil {
		return fmt.Errorf("error pengeluaran 2")
	}

	return nil

}
