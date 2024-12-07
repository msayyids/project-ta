package service

import (
	"context"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/repository"

	"github.com/jmoiron/sqlx"
)

type UserServiceInj interface {
	CreateUser(ctx context.Context, userReq entity.UserRequest) (entity.Users, error)
	FindUserByEmail(ctx context.Context, email string) (entity.Users, error)
	FindUserById(ctx context.Context, id int) (entity.Users, error)
	FindUserByRole(ctx context.Context, role string) (entity.Users, error)
	DeleteUser(ctx context.Context, id int) error
	EditUser(ctx context.Context, id int, userReq entity.UserRequest) (entity.Users, error)
	FindAllUsers(ctx context.Context) ([]entity.Users, error)
}

type UserServices struct {
	DB       *sqlx.DB
	UserRepo repository.UserRepositoryInj
}

func NewUserService(ur repository.UserRepositoryInj, db *sqlx.DB) UserServiceInj {
	return UserServices{
		DB:       db,
		UserRepo: ur,
	}
}

func (s UserServices) CreateUser(ctx context.Context, userReq entity.UserRequest) (entity.Users, error) {

	tx, err := s.DB.Beginx()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	newUser, err := s.UserRepo.AddUser(ctx, userReq, *tx)
	helper.PanicIfError(err)

	return newUser, nil
}

func (s UserServices) FindUserByEmail(ctx context.Context, email string) (entity.Users, error) {
	tx, err := s.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	foundUser, err := s.UserRepo.FindbyEmail(ctx, email, *tx)
	helper.PanicIfError(err)

	return foundUser, nil
}

func (s UserServices) FindUserByRole(ctx context.Context, role string) (entity.Users, error) {
	tx, err := s.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	foundUser, err := s.UserRepo.FindbyRole(ctx, role, *tx)
	helper.PanicIfError(err)

	return foundUser, nil
}

func (s UserServices) FindUserById(ctx context.Context, id int) (entity.Users, error) {
	tx, err := s.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	foundUser, err := s.UserRepo.FindbyId(ctx, id, *tx)
	helper.PanicIfError(err)

	return foundUser, nil
}

func (s UserServices) DeleteUser(ctx context.Context, id int) error {
	tx, err := s.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = s.UserRepo.DeleteUser(ctx, id, *tx)
	helper.PanicIfError(err)

	return nil
}

func (s UserServices) EditUser(ctx context.Context, id int, userReq entity.UserRequest) (entity.Users, error) {
	tx, err := s.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	updatedUser, err := s.UserRepo.EditUser(ctx, id, userReq, *tx)
	helper.PanicIfError(err)

	return updatedUser, nil
}

func (s UserServices) FindAllUsers(ctx context.Context) ([]entity.Users, error) {
	tx, err := s.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users, err := s.UserRepo.FindAllUsers(ctx, *tx)
	helper.PanicIfError(err)

	return users, nil
}
