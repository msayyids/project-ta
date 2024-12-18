package service

import (
	"context"
	"gorm.io/gorm"
	"project-ta/entity"
	"project-ta/repository"
	"time"
)

type KeuntunganService struct {
	DB             *gorm.DB
	KeuntunganRepo repository.KeuntunganRepositoryInj
}

type KeuntunganServiceInj interface {
	GetKeuntunganByDate(ctx context.Context, date time.Time) ([]entity.KeuntunganResponse, error)
	GetKeuntunganByMonth(ctx context.Context, year, month int) ([]entity.KeuntunganResponseMonthly, error)
	GetKeuntunganByLast7Days(ctx context.Context, date time.Time) ([]entity.KeuntunganPer7HariResponse, error)
}

func NewKeuntunganService(db *gorm.DB, repo repository.KeuntunganRepositoryInj) KeuntunganServiceInj {
	return &KeuntunganService{
		DB:             db,
		KeuntunganRepo: repo,
	}
}

// Fungsi untuk mendapatkan keuntungan berdasarkan tanggal tertentu
func (s *KeuntunganService) GetKeuntunganByDate(ctx context.Context, date time.Time) ([]entity.KeuntunganResponse, error) {

	tx := s.DB.Begin()
	keuntunganByDate, err := s.KeuntunganRepo.GetKeuntunganByDate(ctx, tx, date)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return keuntunganByDate, nil
}

// Fungsi untuk mendapatkan keuntungan berdasarkan bulan dan tahun
func (s *KeuntunganService) GetKeuntunganByMonth(ctx context.Context, year, month int) ([]entity.KeuntunganResponseMonthly, error) {
	tx := s.DB.Begin()
	keuntunganByMonth, err := s.KeuntunganRepo.GetKeuntunganByMonth(ctx, tx, year, month)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return keuntunganByMonth, nil
}

// Fungsi untuk mendapatkan keuntungan dalam 7 hari terakhir
func (s *KeuntunganService) GetKeuntunganByLast7Days(ctx context.Context, date time.Time) ([]entity.KeuntunganPer7HariResponse, error) {
	tx := s.DB.Begin()
	keuntunganBy7Day, err := s.KeuntunganRepo.GetKeuntunganByLast7Days(ctx, tx, date)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return keuntunganBy7Day, nil
}
