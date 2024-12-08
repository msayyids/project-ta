package repository

import (
	"context"
	"project-ta/entity"
	"time"

	"gorm.io/gorm"
)

type UserRepositoryInj interface {
	AddUser(ctx context.Context, userReq entity.UserRequest, db *gorm.DB) (entity.Users, error)
	DeleteUser(ctx context.Context, id int, db *gorm.DB) error
	EditUser(ctx context.Context, id int, userReq entity.UserRequest, db *gorm.DB) (entity.Users, error)
	FindbyEmail(ctx context.Context, email string, db *gorm.DB) (entity.Users, error)
	FindbyRole(ctx context.Context, role string, db *gorm.DB) (entity.Users, error)
	FindbyId(ctx context.Context, id int, db *gorm.DB) (entity.Users, error)
	FindAllUsers(ctx context.Context, db *gorm.DB) ([]entity.Users, error)
}

type UserRepositories struct{}

func NewUserRepository() UserRepositoryInj {
	return &UserRepositories{}
}

func (ur *UserRepositories) AddUser(ctx context.Context, userReq entity.UserRequest, db *gorm.DB) (entity.Users, error) {
	user := entity.Users{
		Nama_depan:    userReq.Nama_depan,
		Nama_belakang: userReq.Nama_belakang,
		Role:          userReq.Role,
		Email:         userReq.Email,
		Password:      userReq.Password,
		No_telepon:    userReq.No_telepon,
		Alamat:        userReq.Alamat,
		Gaji:          userReq.Gaji,
		No_rekening:   userReq.No_rekening,
		Bank_id:       userReq.Bank_id,
		Created_At:    time.Now(),
	}

	err := db.Create(&user).Error
	if err != nil {
		return entity.Users{}, err
	}

	return user, nil
}

func (ur *UserRepositories) DeleteUser(ctx context.Context, id int, db *gorm.DB) error {
	err := db.Delete(&entity.Users{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepositories) EditUser(ctx context.Context, id int, userReq entity.UserRequest, db *gorm.DB) (entity.Users, error) {
	var user entity.Users
	err := db.First(&user, id).Error
	if err != nil {
		return entity.Users{}, err
	}

	user.Nama_depan = userReq.Nama_depan
	user.Nama_belakang = userReq.Nama_belakang
	user.Role = userReq.Role
	user.Email = userReq.Email
	user.Password = userReq.Password
	user.No_telepon = userReq.No_telepon
	user.Alamat = userReq.Alamat
	user.Gaji = userReq.Gaji
	user.No_rekening = userReq.No_rekening
	user.Bank_id = userReq.Bank_id
	user.Created_At = time.Now()

	err = db.Save(&user).Error
	if err != nil {
		return entity.Users{}, err
	}

	return user, nil
}

func (ur *UserRepositories) FindbyEmail(ctx context.Context, email string, db *gorm.DB) (entity.Users, error) {
	var user entity.Users
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return entity.Users{}, err
	}
	return user, nil
}

func (ur *UserRepositories) FindbyRole(ctx context.Context, role string, db *gorm.DB) (entity.Users, error) {
	var user entity.Users
	err := db.Where("role = ?", role).First(&user).Error
	if err != nil {
		return entity.Users{}, err
	}
	return user, nil
}

func (ur *UserRepositories) FindbyId(ctx context.Context, id int, db *gorm.DB) (entity.Users, error) {
	var user entity.Users
	err := db.First(&user, id).Error
	if err != nil {
		return entity.Users{}, err
	}
	return user, nil
}

func (ur *UserRepositories) FindAllUsers(ctx context.Context, db *gorm.DB) ([]entity.Users, error) {
	var users []entity.Users
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
