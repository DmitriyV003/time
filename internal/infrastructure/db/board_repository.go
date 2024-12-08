package db

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"tracker/internal/domain/models"
)

type BoardRepository struct {
	db *gorm.DB
}

func NewBoardRepository(db *gorm.DB) *BoardRepository {
	return &BoardRepository{
		db: db,
	}
}

func (r *BoardRepository) Create(board *models.Board) error {
	if board == nil {
		return errors.New("boardRepository.board: board is nil")
	}

	if err := r.db.Model(board).Create(board).Error; err != nil {
		return errors.Wrap(err, "boardRepository.Create: error to create board")
	}

	return nil
}
