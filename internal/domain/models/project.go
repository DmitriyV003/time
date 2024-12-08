package models

import (
	"database/sql"
	"time"
)

type Project struct {
	ID          uint64       `gorm:"id"`
	Name        string       `gorm:"name"`
	WorkspaceID uint64       `gorm:"workspace_id"`
	OwnerID     uint64       `gorm:"owner_id"`
	CreatedAt   time.Time    `gorm:"created_at"`
	UpdatedAt   time.Time    `gorm:"updated_at"`
	DeletedAt   sql.NullTime `gorm:"deleted_at"`
}

func (Project) TableName() string {
	return "projects"
}
