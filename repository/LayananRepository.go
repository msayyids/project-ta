package repository

import (
	"context"
	"project-ta/entity"

	"gorm.io/gorm"
)

type LayananRepositoryInj interface {
	AddLayanan(ctx context.Context, LayananReq entity.LayananRequest, db *gorm.DB) (entity.Layanan, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]entity.Layanan, error)
	EditById(ctx context.Context, id int, db *gorm.DB, layanan entity.LayananRequest) (entity.Layanan, error)
	FindById(ctx context.Context, id int, db *gorm.DB) (entity.Layanan, error)
	DeleteById(ctx context.Context, id int, db *gorm.DB) error
}

type LayananRepository struct{}

func NewLayananRepository() LayananRepositoryInj {
	return LayananRepository{}
}

func (l LayananRepository) AddLayanan(ctx context.Context, LayananReq entity.LayananRequest, db *gorm.DB) (entity.Layanan, error) {
	layanan := entity.Layanan{
		Nama:      LayananReq.Nama,
		Deskripsi: LayananReq.Desksripsi,
		Harga:     LayananReq.Harga,
	}

	err := db.Create(&layanan).Error
	if err != nil {
		return entity.Layanan{}, err
	}

	return layanan, nil
}

func (l LayananRepository) EditById(ctx context.Context, id int, db *gorm.DB, layanan entity.LayananRequest) (entity.Layanan, error) {
	updatedLayanan := entity.Layanan{
		ID:        id,
		Nama:      layanan.Nama,
		Deskripsi: layanan.Desksripsi,
		Harga:     layanan.Harga,
	}

	err := db.Save(&updatedLayanan).Error
	if err != nil {
		return entity.Layanan{}, err
	}

	return updatedLayanan, nil
}

func (l LayananRepository) DeleteById(ctx context.Context, id int, db *gorm.DB) error {
	var layanan entity.Layanan
	err := db.Delete(&layanan, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (l LayananRepository) FindById(ctx context.Context, id int, db *gorm.DB) (entity.Layanan, error) {
	var layanan entity.Layanan
	err := db.First(&layanan, id).Error
	if err != nil {
		return entity.Layanan{}, err
	}

	return layanan, nil
}

func (l LayananRepository) FindAll(ctx context.Context, db *gorm.DB) ([]entity.Layanan, error) {
	var layanan []entity.Layanan
	err := db.Find(&layanan).Error
	if err != nil {
		return nil, err
	}

	return layanan, nil
}
