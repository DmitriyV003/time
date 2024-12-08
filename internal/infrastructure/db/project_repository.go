package db

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"tracker/internal/domain/models"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

func (r *ProjectRepository) Create(project *models.Project) error {
	if project == nil {
		return errors.New("repository.Create: project is nil")
	}

	if err := r.db.Model(project).Create(project).Error; err != nil {
		return errors.Wrap(err, "repository.Create: error to create project")
	}

	return nil
}
