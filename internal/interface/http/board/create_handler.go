package board

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"tracker/internal/application/dto"
	"tracker/internal/application/service"
	"tracker/internal/domain/models"
)

type request struct {
	WorkspaceID uint64  `json:"workspace_id" validate:"required"`
	ProjectID   *uint64 `json:"project_id"`
	Name        string  `json:"name" validate:"required"`
}

type CreateHandler struct {
	boardService *service.BoardService
}

func NewCreateHandler(boardService *service.BoardService) *CreateHandler {
	return &CreateHandler{boardService: boardService}
}

func (h *CreateHandler) Handle(c echo.Context) error {
	var (
		req   request
		userr *models.User
	)

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}

	par := dto.CreateBoardDTO{
		WorkspaceID: req.WorkspaceID,
		Name:        req.Name,
	}

	if req.ProjectID != nil {
		par.ProjectID = sql.Null[uint64]{Valid: true, V: *req.ProjectID}
	}

	_, err = h.boardService.Create(userr, par)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error to create board")
	}

	return c.JSON(http.StatusCreated, nil)
}
