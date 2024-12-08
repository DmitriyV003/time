package db

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"tracker/internal/domain/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *models.User) error {
	if user == nil {
		return errors.New("userRepository.Save: user is nil")
	}

	if err := r.db.Model(user).Create(user).Error; err != nil {
		return errors.Wrap(err, "userRepository.Save: error to create user")
	}

	return nil
}
