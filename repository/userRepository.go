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
	DeleteUser(ctx context.Context, id int, tx sqlx.Tx) error
	EditUser(ctx context.Context, id int, userReq entity.UserRequest, tx sqlx.Tx) (entity.Users, error)
	FindbyEmail(ctx context.Context, email string, tx sqlx.Tx) (entity.Users, error)
	FindbyRole(ctx context.Context, role string, tx sqlx.Tx) (entity.Users, error)
	FindbyId(ctx context.Context, id int, tx sqlx.Tx) (entity.Users, error)
	FindAllUsers(ctx context.Context, tx sqlx.Tx) ([]entity.Users, error)
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

func (ur UserRepositories) DeleteUser(ctx context.Context, id int, tx sqlx.Tx) error {
	sqlQuery := `DELETE FROM users WHERE id = $1`

	// Execute the delete query
	_, err := tx.ExecContext(ctx, sqlQuery, id)
	if err != nil {
		return err
	}

	return nil
}

func (ur UserRepositories) EditUser(ctx context.Context, id int, userReq entity.UserRequest, tx sqlx.Tx) (entity.Users, error) {
	sqlQuery := `
        UPDATE users 
        SET nama_depan = $1, nama_belakang = $2, role = $3, email = $4, 
            no_telepon = $5, alamat = $6, gaji = $7, no_rekening = $8, bank_id = $9, created_at = $10
        WHERE id = $11
        RETURNING *
    `

	var updatedUser entity.Users
	err := tx.QueryRowxContext(ctx, sqlQuery,
		userReq.Nama_depan,
		userReq.Nama_belakang,
		userReq.Role,
		userReq.Email,
		userReq.No_telepon,
		userReq.Alamat,
		userReq.Gaji,
		userReq.No_rekening,
		userReq.Bank_id,
		time.Now(), // Update timestamp
		id,         // User ID for the update
	).StructScan(&updatedUser)

	if err != nil {
		return entity.Users{}, err
	}

	return updatedUser, nil
}

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

func (ur UserRepositories) FindAllUsers(ctx context.Context, tx sqlx.Tx) ([]entity.Users, error) {
	sqlQuery := `SELECT * FROM users`

	var users []entity.Users

	err := tx.SelectContext(ctx, &users, sqlQuery)
	if err != nil {
		return nil, err
	}

	return users, nil
}
