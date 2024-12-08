package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint64         `gorm:"id;primaryKey"`
	Email     string         `gorm:"email"`
	Name      string         `gorm:"name"`
	LastName  sql.NullString `gorm:"last_name"`
	Password  string         `gorm:"password"`
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at"`
	DeletedAt sql.NullTime   `gorm:"deleted_at"`
}

func (User) TableName() string {
	return "users"
}
