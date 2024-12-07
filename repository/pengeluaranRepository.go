package repository

import (
	"context"
	"fmt"
	"project-ta/entity"
	"project-ta/helper"
	"time"

	"github.com/jmoiron/sqlx"
)

type PengeluaranRepositoryInj interface {
	AddPengeluaran(ctx context.Context, tx sqlx.Tx, pengeluaran entity.PengeluaranRequest) (entity.Pengeluaran, error)
	FindAllPengeluaran(ctx context.Context, tx sqlx.Tx) ([]entity.Pengeluaran, error)
	FindPengeluaranById(ctx context.Context, tx sqlx.Tx, id int) (entity.Pengeluaran, error)
	DeletePengeluaran(ctx context.Context, tx sqlx.Tx, id int) error
	UpdatePengeluaran(ctx context.Context, tx sqlx.Tx, id int, pengeluaran entity.PengeluaranRequest) (entity.Pengeluaran, error)
}

type pengeluaranRepository struct {
}

func NewPengeluaranRepository() PengeluaranRepositoryInj {
	return &pengeluaranRepository{}
}

func (r *pengeluaranRepository) AddPengeluaran(ctx context.Context, tx sqlx.Tx, pengeluaran entity.PengeluaranRequest) (entity.Pengeluaran, error) {
	query := `
		INSERT INTO pengeluaran (nama_pengeluaran, keterangan, users_id,total, bukti_pengeluaran,created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *
	`

	var newPengeluaran entity.Pengeluaran
	err := tx.QueryRowxContext(ctx, query,
		pengeluaran.Nama_pengeluaran, pengeluaran.Keterangan, pengeluaran.Users_id, pengeluaran.Total, pengeluaran.Bukti_pengeluaran, time.Now(),
	).StructScan(&newPengeluaran)
	// helper.PanicIfError(err)
	if err != nil {
		return entity.Pengeluaran{}, fmt.Errorf("error pengeluaran 1")
	}

	return newPengeluaran, nil
}

func (r *pengeluaranRepository) FindAllPengeluaran(ctx context.Context, tx sqlx.Tx) ([]entity.Pengeluaran, error) {
	query := `SELECT * FROM pengeluaran
	`

	var pengeluaran []entity.Pengeluaran

	err := tx.SelectContext(ctx, &pengeluaran, query)
	if err != nil {
		return []entity.Pengeluaran{}, err
	}

	return pengeluaran, nil
}

func (r *pengeluaranRepository) FindPengeluaranById(ctx context.Context, tx sqlx.Tx, id int) (entity.Pengeluaran, error) {
	query := `SELECT * FROM pengeluaran WHERE id = $1
	`

	var pengeluaran entity.Pengeluaran

	err := tx.QueryRowxContext(ctx, query, id).StructScan(&pengeluaran)

	if err != nil {
		return entity.Pengeluaran{}, err
	}

	return pengeluaran, nil
}

func (r *pengeluaranRepository) DeletePengeluaran(ctx context.Context, tx sqlx.Tx, id int) error {
	sqlQuery := `DELETE FROM pengeluaran WHERE id = $1`

	// Execute the delete query
	_, err := tx.ExecContext(ctx, sqlQuery, id)
	helper.PanicIfError(err)

	return nil
}

func (r *pengeluaranRepository) UpdatePengeluaran(ctx context.Context, tx sqlx.Tx, id int, pengeluaran entity.PengeluaranRequest) (entity.Pengeluaran, error) {
	query := `
		UPDATE pengeluaran 
		SET 
			nama_pengeluaran = $1, 
			keterangan = $2, 
			users_id = $3, 
			total = $4, 
			bukti_pengeluaran = $5, 
			created_at = $6
		WHERE id = $7
		RETURNING *
	`

	var updatedPengeluaran entity.Pengeluaran
	err := tx.QueryRowxContext(ctx, query,
		pengeluaran.Nama_pengeluaran, pengeluaran.Keterangan, pengeluaran.Users_id, pengeluaran.Total, pengeluaran.Bukti_pengeluaran, time.Now(), id,
	).StructScan(&updatedPengeluaran)
	if err != nil {
		return entity.Pengeluaran{}, fmt.Errorf("error updating pengeluaran: %w", err)
	}

	return updatedPengeluaran, nil
}
