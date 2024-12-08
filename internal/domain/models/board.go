package models

import (
	"database/sql"
	"time"
)

type Board struct {
	Id          uint64           `gorm:"id"`
	Name        string           `gorm:"name"`
	WorkspaceID uint64           `gorm:"workspace_id"`
	OwnerID     uint64           `gorm:"owner_id"`
	ProjectID   sql.Null[uint64] `gorm:"project_id"`
	CreatedAt   time.Time        `gorm:"created_at"`
	UpdatedAt   time.Time        `gorm:"updated_at"`
	DeletedAt   sql.NullTime     `gorm:"deleted_at"`
}

func (Board) TableName() string {
	return "boards"
}
