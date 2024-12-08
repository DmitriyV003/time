package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
	"tracker/internal/application/dto"
	"tracker/internal/domain/models"
	"tracker/internal/infrastructure/db"
)

type UserService struct {
	userRepository *db.UserRepository
}

func (s *UserService) Create(dto dto.CreateUserDTO) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "error to hash password")
	}
	user := models.User{
		Email:     dto.Email,
		Name:      dto.Name,
		Password:  string(hash),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: sql.NullTime{},
	}

	if dto.LastName != "" {
		user.LastName = sql.NullString{
			String: dto.LastName,
			Valid:  true,
		}
	}
	err = s.userRepository.Create(&user)
	if err != nil {
		return nil, errors.Wrap(err, "userService.create: error to create user")
	}

	return &user, nil
}
