package service

import (
	"context"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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
	DB        *gorm.DB
	UserRepo  repository.UserRepositoryInj
	Validator validator.Validate
}

func NewUserService(ur repository.UserRepositoryInj, db *gorm.DB, v validator.Validate) UserServiceInj {
	return &UserServices{
		DB:        db,
		UserRepo:  ur,
		Validator: v,
	}
}

func (s *UserServices) CreateUser(ctx context.Context, userReq entity.UserRequest) (entity.Users, error) {
	err := s.Validator.Struct(userReq)
	if err != nil {
		return entity.Users{}, err
	}

	tx := s.DB.Begin()

	defer helper.CommitOrRollback(tx)

	newUser, err := s.UserRepo.AddUser(ctx, userReq, tx)
	if err != nil {
		return entity.Users{}, err
	}

	return newUser, nil
}

func (s *UserServices) FindUserByEmail(ctx context.Context, email string) (entity.Users, error) {
	var user entity.Users
	tx := s.DB.Begin()

	defer helper.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return entity.Users{}, err
	}

	return user, nil
}

func (s *UserServices) FindUserByRole(ctx context.Context, role string) (entity.Users, error) {
	var user entity.Users
	tx := s.DB.Begin()

	defer helper.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Where("role = ?", role).First(&user).Error
	if err != nil {
		return entity.Users{}, err
	}

	return user, nil
}

func (s *UserServices) FindUserById(ctx context.Context, id int) (entity.Users, error) {
	var user entity.Users
	tx := s.DB.Begin()

	defer helper.CommitOrRollback(tx)

	err := tx.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return entity.Users{}, err
	}

	return user, nil
}

func (s *UserServices) DeleteUser(ctx context.Context, id int) error {
	tx := s.DB.Begin()

	defer helper.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Delete(&entity.Users{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *UserServices) EditUser(ctx context.Context, id int, userReq entity.UserRequest) (entity.Users, error) {
	var user entity.Users
	tx := s.DB.Begin()

	defer helper.CommitOrRollback(tx)

	err := tx.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return entity.Users{}, err
	}

	// user.Name = userReq.Name
	// user.Email = userReq.Email
	// Update other fields as necessary

	err = tx.WithContext(ctx).Save(&user).Error
	if err != nil {
		return entity.Users{}, err
	}

	return user, nil
}

func (s *UserServices) FindAllUsers(ctx context.Context) ([]entity.Users, error) {
	var users []entity.Users
	tx := s.DB.Begin()

	defer helper.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
