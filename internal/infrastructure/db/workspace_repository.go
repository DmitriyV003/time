package db

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"tracker/internal/domain/models"
)

type WorkspaceRepository struct {
	db *gorm.DB
}

func NewWorkspaceRepository(db *gorm.DB) *WorkspaceRepository {
	return &WorkspaceRepository{
		db: db,
	}
}

func (r *WorkspaceRepository) Create(workspace *models.Workspace) error {
	if workspace == nil {
		return errors.New("workspaceRepository.workspace: workspace is nil")
	}

	if err := r.db.Model(workspace).Create(workspace).Error; err != nil {
		return errors.Wrap(err, "workspaceRepository.Save: error to create workspace")
	}

	return nil
}
