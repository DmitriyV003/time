package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"time"
	"tracker/internal/application/dto"
	"tracker/internal/domain/models"
	db2 "tracker/internal/infrastructure/db"
)

type WorkspaceService struct {
	workspaceRepository *db2.WorkspaceRepository
}

func NewWorkspaceService(db *gorm.DB) *WorkspaceService {
	return &WorkspaceService{
		workspaceRepository: db2.NewWorkspaceRepository(db),
	}
}

func (s *WorkspaceService) Create(ctx context.Context, user *models.User, dto dto.CreateWorkspaceDTO) (*models.Workspace, error) {
	//if user == nil {
	//	return nil, errors.New("error to create workspace, user in nil")
	//}
	ws := models.Workspace{
		//OwnerID:   user.ID,
		OwnerID:   1,
		Name:      dto.Name,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: sql.NullTime{},
	}

	err := s.workspaceRepository.Create(&ws)
	if err != nil {
		return nil, errors.Wrap(err, "workspaceService.create: error to create workspace")
	}

	return &ws, nil
}
