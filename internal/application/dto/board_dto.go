package dto

import "database/sql"

type CreateBoardDTO struct {
	WorkspaceID uint64
	ProjectID   sql.Null[uint64]
	Name        string
}
