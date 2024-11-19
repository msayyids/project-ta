package repository

import (
	"context"
	"project-ta/entity"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepositories struct{}

type UserRepositoryInj interface {
	AddUser(ctx context.Context, userReq entity.UserRequest, db sqlx.Tx) (entity.Users, error)
	DeleteUser()
	EditUser()
	FindbyEmail(ctx context.Context, email string, tx sqlx.Tx) (entity.Users, error)
	FindbyRole(ctx context.Context, role string, tx sqlx.Tx) (entity.Users, error)
	FindbyId(ctx context.Context, id int, tx sqlx.Tx) (entity.Users, error)
}

func NewUserRepository() UserRepositoryInj {
	return UserRepositories{}
}

func (ur UserRepositories) AddUser(ctx context.Context, userReq entity.UserRequest, tx sqlx.Tx) (entity.Users, error) {
	sqlQuery := `
        INSERT INTO users (
            nama_depan, nama_belakang, role, email, password, no_telepon, alamat, gaji, no_rekening, bank_id, created_at
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
        ) RETURNING *
    `
	var newUser entity.Users
	err := tx.QueryRowxContext(ctx, sqlQuery,
		userReq.Nama_depan,
		userReq.Nama_belakang,
		userReq.Role,
		userReq.Email,
		userReq.Password,
		userReq.No_telepon,
		userReq.Alamat,
		userReq.Gaji,
		userReq.No_rekening,
		userReq.Bank_id,
		time.Now(),
	).StructScan(&newUser)

	if err != nil {
		return entity.Users{}, err
	}

	return newUser, nil
}

func (ur UserRepositories) DeleteUser() {}

func (ur UserRepositories) EditUser() {}

func (ur UserRepositories) FindbyEmail(ctx context.Context, email string, tx sqlx.Tx) (entity.Users, error) {
	sqlQuery := `SELECT * FROM users WHERE email = $1`

	var user entity.Users

	err := tx.QueryRowxContext(ctx, sqlQuery, email).StructScan(&user)

	if err != nil {
		return entity.Users{}, err
	}

	return user, nil
}

func (ur UserRepositories) FindbyRole(ctx context.Context, role string, tx sqlx.Tx) (entity.Users, error) {
	sqlQuery := `SELECT * FROM users WHERE role = $1`

	var user entity.Users

	err := tx.QueryRowxContext(ctx, sqlQuery, role).StructScan(&user)

	if err != nil {
		return entity.Users{}, err
	}

	return user, nil
}

func (ur UserRepositories) FindbyId(ctx context.Context, id int, tx sqlx.Tx) (entity.Users, error) {
	sqlQuery := `SELECT * FROM users WHERE id = $1`

	var user entity.Users

	err := tx.QueryRowxContext(ctx, sqlQuery, id).StructScan(&user)

	if err != nil {
		return entity.Users{}, err
	}

	return user, nil
}
