package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
	"tracker/internal/application/dto"
	"tracker/internal/domain/models"
	db2 "tracker/internal/infrastructure/db"
)

type ProjectService struct {
	projectRepository *db2.ProjectRepository
}

func NewProjectService(db *gorm.DB) *ProjectService {
	return &ProjectService{
		projectRepository: db2.NewProjectRepository(db),
	}
}

func (s *ProjectService) Create(user *models.User, dto dto.CreateBoardDTO) (*models.Project, error) {
	if user == nil {
		return nil, errors.New("error to create project, user is nil")
	}

	project := models.Project{
		Name:        dto.Name,
		WorkspaceID: dto.WorkspaceID,
		OwnerID:     user.ID,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   sql.NullTime{},
	}

	err := s.projectRepository.Create(&project)
	if err != nil {
		return nil, errors.Wrap(err, "service.create: error to create project")
	}

	return &project, nil
}
