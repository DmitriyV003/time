package models

import (
	"database/sql"
	"time"
)

type BoardSection struct {
	ID        uint64       `gorm:"id"`
	BoardID   uint64       `gorm:"board_id"`
	Name      string       `gorm:"name"`
	CreatedAt time.Time    `gorm:"created_at"`
	UpdatedAt time.Time    `gorm:"updated_at"`
	DeletedAt sql.NullTime `gorm:"deleted_at"`
}

func (BoardSection) TableName() string {
	return "board_sections"
}
