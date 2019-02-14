package main

import (
	pb "github.com/ashurai/fap-back/user-service/proto/user"
	"github.com/jinzhu/gorm"
)

// Repository to user manage user interface
type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (pb.User, error)
	Create(user *pb.User) error
	GetByEmailAndPassword(user *pb.User) (*pb.User, error)
}

// UserRepository struct
type UserRepository struct {
	db *gorm.DB
}

// GetAll Users from user object
func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Get User by ID
func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetByEmailAndPassword to validate users
func (repo *UserRepository) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	if repo.db.First(&user).Error; err != nil {
		return nil, error
	}
	return user, nil
}

// Create user / register user
func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
}
