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
	FindUSerByemail(ctx context.Context, email string) (entity.Users, error)
	FindUSerById(ctx context.Context, id int) (entity.Users, error)
	FindUSerByRole(ctx context.Context, role string) (entity.Users, error)
}

type UserServices struct {
	DB       *sqlx.DB
	UserRepo repository.UserRepositoryInj
}

func NewUserService(ur repository.UserRepositoryInj, db *sqlx.DB) UserServiceInj {
	return UserServices{
		DB:       db,
		UserRepo: ur}
}

func (s UserServices) CreateUser(ctx context.Context, userReq entity.UserRequest) (entity.Users, error) {
	tx, err := s.DB.Beginx()
	// helper.PanicIfError(err)
	if err != nil {
		return entity.Users{}, err
	}
	defer helper.CommitOrRollback(tx)

	newUsers, err := s.UserRepo.AddUser(ctx, userReq, *tx)
	helper.PanicIfError(err)

	return newUsers, nil
}

func (s UserServices) FindUSerByemail(ctx context.Context, email string) (entity.Users, error) {
	tx, err := s.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	foundUser, err := s.UserRepo.FindbyEmail(ctx, email, *tx)
	helper.PanicIfError(err)

	return foundUser, nil
}

func (s UserServices) FindUSerByRole(ctx context.Context, role string) (entity.Users, error) {
	tx, err := s.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	foundUser, err := s.UserRepo.FindbyRole(ctx, role, *tx)
	helper.PanicIfError(err)

	return foundUser, nil
}

func (s UserServices) FindUSerById(ctx context.Context, id int) (entity.Users, error) {
	tx, err := s.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	foundUser, err := s.UserRepo.FindbyId(ctx, id, *tx)
	helper.PanicIfError(err)

	return foundUser, nil
}
