package project

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tracker/internal/application/dto"
	"tracker/internal/application/service"
	"tracker/internal/domain/models"
)

type request struct {
	Name        string `json:"name" validate:"required"`
	WorkspaceID uint64 `json:"workspace_id" validate:"required"`
}

type CreateHandler struct {
	projectService *service.ProjectService
}

func NewCreateHandler(projectService *service.ProjectService) *CreateHandler {
	return &CreateHandler{projectService: projectService}
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
		Name:        req.Name,
		WorkspaceID: req.WorkspaceID,
	}

	_, err = h.projectService.Create(userr, par)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error to create project")
	}

	return c.JSON(http.StatusCreated, nil)
}
