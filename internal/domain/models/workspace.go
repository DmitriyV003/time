package models

import (
	"database/sql"
	"time"
)

type Workspace struct {
	ID        uint64       `gorm:"id"`
	Name      string       `gorm:"string"`
	OwnerID   uint64       `gorm:"owner_id"`
	CreatedAt time.Time    `gorm:"created_at"`
	UpdatedAt time.Time    `gorm:"updated_at"`
	DeletedAt sql.NullTime `gorm:"deleted_at"`
}

func (Workspace) TableName() string {
	return "workspaces"
}
