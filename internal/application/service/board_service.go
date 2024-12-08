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

type BoardService struct {
	boardRepository *db2.BoardRepository
}

func NewBoardService(db *gorm.DB) *BoardService {
	return &BoardService{
		boardRepository: db2.NewBoardRepository(db),
	}
}

func (s *BoardService) Create(user *models.User, dto dto.CreateBoardDTO) (*models.Board, error) {
	if user == nil {
		return nil, errors.New("error to create board, user is nil")
	}

	board := models.Board{
		WorkspaceID: dto.WorkspaceID,
		OwnerID:     user.ID,
		ProjectID:   dto.ProjectID,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   sql.NullTime{},
	}

	err := s.boardRepository.Create(&board)
	if err != nil {
		return nil, errors.Wrap(err, "boardService.create: error to create board")
	}

	return &board, nil
}
