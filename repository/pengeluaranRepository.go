package repository

import (
	"context"
	"fmt"
	"project-ta/entity"

	"gorm.io/gorm"
)

type PengeluaranRepositoryInj interface {
	AddPengeluaran(ctx context.Context, pengeluaran entity.PengeluaranRequest, db *gorm.DB) (entity.Pengeluaran, error)
	FindAllPengeluaran(ctx context.Context, db *gorm.DB) ([]entity.Pengeluaran, error)
	FindPengeluaranById(ctx context.Context, id int, db *gorm.DB) (entity.Pengeluaran, error)
	DeletePengeluaran(ctx context.Context, id int, db *gorm.DB) error
	UpdatePengeluaran(ctx context.Context, id int, pengeluaran entity.PengeluaranRequest, db *gorm.DB) (entity.Pengeluaran, error)
}

type pengeluaranRepository struct {
}

func NewPengeluaranRepository() PengeluaranRepositoryInj {
	return &pengeluaranRepository{}
}

func (r *pengeluaranRepository) AddPengeluaran(ctx context.Context, pengeluaran entity.PengeluaranRequest, db *gorm.DB) (entity.Pengeluaran, error) {
	var newPengeluaran entity.Pengeluaran
	// Create menggunakan GORM
	err := db.WithContext(ctx).Create(&newPengeluaran).Error
	if err != nil {
		return entity.Pengeluaran{}, fmt.Errorf("error creating pengeluaran: %w", err)
	}
	return newPengeluaran, nil
}

func (r *pengeluaranRepository) FindAllPengeluaran(ctx context.Context, db *gorm.DB) ([]entity.Pengeluaran, error) {
	var pengeluaran []entity.Pengeluaran
	// FindAll menggunakan GORM
	err := db.WithContext(ctx).Find(&pengeluaran).Error
	if err != nil {
		return nil, fmt.Errorf("error fetching all pengeluaran: %w", err)
	}
	return pengeluaran, nil
}

func (r *pengeluaranRepository) FindPengeluaranById(ctx context.Context, id int, db *gorm.DB) (entity.Pengeluaran, error) {
	var pengeluaran entity.Pengeluaran
	// Find by ID menggunakan GORM
	err := db.WithContext(ctx).First(&pengeluaran, id).Error
	if err != nil {
		return entity.Pengeluaran{}, fmt.Errorf("error finding pengeluaran by id: %w", err)
	}
	return pengeluaran, nil
}

func (r *pengeluaranRepository) DeletePengeluaran(ctx context.Context, id int, db *gorm.DB) error {
	// Delete menggunakan GORM
	err := db.WithContext(ctx).Delete(&entity.Pengeluaran{}, id).Error
	if err != nil {
		return fmt.Errorf("error deleting pengeluaran: %w", err)
	}
	return nil
}

func (r *pengeluaranRepository) UpdatePengeluaran(ctx context.Context, id int, pengeluaran entity.PengeluaranRequest, db *gorm.DB) (entity.Pengeluaran, error) {
	var updatedPengeluaran entity.Pengeluaran
	// Update menggunakan GORM
	err := db.WithContext(ctx).Model(&updatedPengeluaran).Where("id = ?", id).Updates(pengeluaran).Error
	if err != nil {
		return entity.Pengeluaran{}, fmt.Errorf("error updating pengeluaran: %w", err)
	}
	return updatedPengeluaran, nil
}
