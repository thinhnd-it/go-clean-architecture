package repository

import (
	"context"
	"errors"
	contract "go-clean-architecture/module/user/contract/repository"
	"go-clean-architecture/module/user/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) contract.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(c context.Context, user *model.User) error {
	if err := r.db.WithContext(c).Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Fetch(c context.Context) ([]model.User, error) {
	var users []model.User

	if err := r.db.WithContext(c).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetByUsernameOrEmail(c context.Context, usernameOrEmail string) (model.User, error) {
	var user model.User

	if err := r.db.WithContext(c).Where("email = ?", usernameOrEmail).Or("username = ?", usernameOrEmail).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, nil
		}
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetByEmail(c context.Context, email string) (model.User, error) {
	var user model.User

	if err := r.db.WithContext(c).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, nil
		}
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetByID(c context.Context, id string) (model.User, error) {
	var user model.User

	if err := r.db.WithContext(c).Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, nil
		}
		return model.User{}, err
	}

	return user, nil
}
