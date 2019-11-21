package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/phuonghau98/walkie3/srv/user/proto/user"
)

type Repository interface {
	Create (user *pb.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) Create(user *pb.User) error {
	userObj := &User{
		Email: user.Email,
		Password: user.Password,
		Firstname: user.FirstName,
		Lastname: user.LastName,
	}
	if err := repo.db.Create(userObj).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetByEmail (email string) (*User, error) {
	user := &User{}
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByID (id string) (*User, error) {
	user := &User{}
	if err := repo.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}