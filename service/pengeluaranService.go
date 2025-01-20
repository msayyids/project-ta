package service

import (
	"context"
	"fmt"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/repository"
	"time"

	"gorm.io/gorm"
)

type PengeluaranService struct {
	DB              *gorm.DB
	PengeluaranRepo repository.PengeluaranRepositoryInj
}

type PengeluaranServiceInj interface {
	CreatePengeluaran(ctx context.Context, pengeluaranReq entity.PengeluaranRequest) (entity.Pengeluaran, error)
	EditPengeluaran(ctx context.Context, pengeluaranReq entity.PengeluaranRequest, id int) (entity.Pengeluaran, error)
	FindPengeluaranById(ctx context.Context, id int) (entity.Pengeluaran, error)
	FindAllPengeluaran(ctx context.Context) ([]entity.Pengeluaran, error)
	DeletePengeluaran(ctx context.Context, id int) error
	GetPengeluaranByDate(ctx context.Context, date time.Time) ([]entity.Pengeluaran, error)
}

func NewPengeluaranService(db *gorm.DB, pr repository.PengeluaranRepositoryInj) PengeluaranServiceInj {
	return &PengeluaranService{
		DB:              db,
		PengeluaranRepo: pr,
	}
}

func (pr *PengeluaranService) GetPengeluaranByDate(ctx context.Context, date time.Time) ([]entity.Pengeluaran, error) {

	tx := pr.DB.Begin()

	defer helper.CommitOrRollback(tx)

	pengeluarans, err := pr.PengeluaranRepo.FindByDateRange(ctx, date, tx)
	if err != nil {
		return []entity.Pengeluaran{}, err
	}

	return pengeluarans, nil
}

func (pr *PengeluaranService) CreatePengeluaran(ctx context.Context, pengeluaranReq entity.PengeluaranRequest) (entity.Pengeluaran, error) {
	tx := pr.DB.Begin()
	defer helper.CommitOrRollback(tx)

	newPengeluaran, err := pr.PengeluaranRepo.AddPengeluaran(ctx, pengeluaranReq, tx)
	if err != nil {
		return entity.Pengeluaran{}, fmt.Errorf("error saat menambahkan pengeluaran: %w", err)
	}

	return newPengeluaran, nil
}

func (pr *PengeluaranService) EditPengeluaran(ctx context.Context, pengeluaranReq entity.PengeluaranRequest, id int) (entity.Pengeluaran, error) {
	tx := pr.DB.Begin()
	defer helper.CommitOrRollback(tx)

	// Panggil repository untuk update pengeluaran
	updatedPengeluaran, err := pr.PengeluaranRepo.UpdatePengeluaran(ctx, id, pengeluaranReq, tx)
	if err != nil {
		return entity.Pengeluaran{}, fmt.Errorf("error saat mengedit pengeluaran: %w", err)
	}

	return updatedPengeluaran, nil
}

func (pr *PengeluaranService) FindPengeluaranById(ctx context.Context, id int) (entity.Pengeluaran, error) {
	tx := pr.DB.Begin()
	defer helper.CommitOrRollback(tx)

	pengeluaran, err := pr.PengeluaranRepo.FindPengeluaranById(ctx, id, tx)
	if err != nil {
		return entity.Pengeluaran{}, fmt.Errorf("error saat mencari pengeluaran: %w", err)
	}

	return pengeluaran, nil
}

func (pr *PengeluaranService) FindAllPengeluaran(ctx context.Context) ([]entity.Pengeluaran, error) {
	tx := pr.DB.Begin()
	defer helper.CommitOrRollback(tx)

	pengeluarans, err := pr.PengeluaranRepo.FindAllPengeluaran(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("error saat mencari semua pengeluaran: %w", err)
	}

	return pengeluarans, nil
}

func (pr *PengeluaranService) DeletePengeluaran(ctx context.Context, id int) error {
	tx := pr.DB.Begin()
	defer helper.CommitOrRollback(tx)

	err := pr.PengeluaranRepo.DeletePengeluaran(ctx, id, tx)
	if err != nil {
		return fmt.Errorf("error saat menghapus pengeluaran: %w", err)
	}

	return nil
}
