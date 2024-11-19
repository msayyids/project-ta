package repository

import (
	"context"
	"project-ta/entity"
	"project-ta/helper"

	"github.com/jmoiron/sqlx"
)

type LayananRepositories struct{}

type LayananRepositoryInj interface {
	AddLayanan(ctx context.Context, LayananReq entity.LayananRequest, db sqlx.Tx) (entity.Layanan, error)
	FindAll(ctx context.Context, db sqlx.Tx) ([]entity.Layanan, error)
	EditById(ctx context.Context, id int, db *sqlx.Tx, layanan entity.LayananRequest) (entity.Layanan, error)
	FindById(ctx context.Context, id int, db sqlx.Tx) (entity.Layanan, error)
	DeleteById(ctx context.Context, id int, db *sqlx.Tx) error
}

func NewLayananRepository() LayananRepositoryInj {
	return LayananRepositories{}
}

func (l LayananRepositories) AddLayanan(ctx context.Context, LayananReq entity.LayananRequest, db sqlx.Tx) (entity.Layanan, error) {
	sqlQuery := `INSERT INTO layanan (nama,deskripsi) VALUES ($1,$2) RETURNING *`

	var layanan entity.Layanan

	err := db.QueryRowxContext(ctx, sqlQuery, LayananReq.Nama, LayananReq.Harga, LayananReq.Desksripsi).StructScan(&layanan)
	helper.PanicIfError(err)

	return layanan, nil
}

func (l LayananRepositories) EditById(ctx context.Context, id int, db *sqlx.Tx, layanan entity.LayananRequest) (entity.Layanan, error) {
	sqlQuery := `UPDATE layanan
                 SET nama = $1, deskripsi = $2, harga = $3
                 WHERE id = $4
                 RETURNING id, nama, deskripsi, harga;`

	var updatedLayanan entity.Layanan
	err := db.GetContext(ctx, &updatedLayanan, sqlQuery, layanan.Nama, layanan.Desksripsi, layanan.Harga, id)
	if err != nil {
		return entity.Layanan{}, err
	}

	return updatedLayanan, nil
}

func (l LayananRepositories) DeleteById(ctx context.Context, id int, db *sqlx.Tx) error {
	sqlQuery := `DELETE FROM layanan WHERE id = $1;`

	_, err := db.ExecContext(ctx, sqlQuery, id)
	if err != nil {
		return err
	}

	return nil
}

func (l LayananRepositories) FindById(ctx context.Context, id int, db sqlx.Tx) (entity.Layanan, error) {
	sqlQuery := `SELECT * FROM layanan where id = $1`

	var layananById entity.Layanan

	err := db.QueryRowxContext(ctx, sqlQuery, id).StructScan(&layananById)
	helper.PanicIfError(err)
	return layananById, nil
}

func (l LayananRepositories) FindAll(ctx context.Context, db sqlx.Tx) ([]entity.Layanan, error) {
	sqlQuery := `SELECT * FROM layanan`

	var layanan []entity.Layanan

	err := db.SelectContext(ctx, &layanan, sqlQuery)
	helper.PanicIfError(err)

	return layanan, nil
}
