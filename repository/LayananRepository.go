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
	var existingLayanan entity.Layanan

	// Cari layanan berdasarkan ID
	if err := db.WithContext(ctx).First(&existingLayanan, id).Error; err != nil {
		return entity.Layanan{}, err
	}

	// Update kolom-kolom yang disertakan dalam request
	if layanan.Nama != "" {
		existingLayanan.Nama = layanan.Nama
	}
	if layanan.Desksripsi != "" {
		existingLayanan.Deskripsi = layanan.Desksripsi
	}
	if layanan.Harga != 0 {
		existingLayanan.Harga = layanan.Harga
	}

	// Simpan perubahan ke dalam database
	err := db.WithContext(ctx).Save(&existingLayanan).Error
	if err != nil {
		return entity.Layanan{}, err
	}

	return existingLayanan, nil
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
