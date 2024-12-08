package workspace

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
	"tracker/internal/application/dto"
	"tracker/internal/application/service"
	"tracker/internal/application/translator"
	"tracker/internal/domain/models"
	"tracker/internal/interface/http/api_response"
)

type request struct {
	Name string `json:"name" validate:"required"`
}

type CreateHandler struct {
	workspaceService *service.WorkspaceService
	translator       *translator.Translator
	validator        *validator.Validate
}

func NewCreateHandler(
	workspaceService *service.WorkspaceService,
	translator *translator.Translator,
	validator *validator.Validate,
) *CreateHandler {
	return &CreateHandler{
		workspaceService: workspaceService,
		translator:       translator,
		validator:        validator,
	}
}

func (h *CreateHandler) Handle(c echo.Context) error {
	var (
		req            request
		userr          *models.User
		errs           validator.ValidationErrors
		responseErrors map[string]string
		ws             *models.Workspace
	)

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, api_response.CreateErrorApiResponse(err.Error()))
	}

	err = h.validator.Struct(req)
	ok := errors.As(err, &errs)
	if ok {
		responseErrors = h.translator.TranslateError(err)
		return c.JSON(http.StatusBadRequest, responseErrors)
	}
	if err != nil && !ok {
		return c.JSON(http.StatusUnprocessableEntity, api_response.CreateErrorApiResponse("error to create validation errors"))
	}

	par := dto.CreateWorkspaceDTO{
		Name: req.Name,
	}
	ws, err = h.workspaceService.Create(c.Request().Context(), userr, par)
	if err != nil {
		log.Err(err).Msg("error to create workspace")
		return c.JSON(http.StatusUnprocessableEntity, api_response.CreateErrorApiResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, mapToWorkspaceResponse(ws))
}

type Response struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func mapToWorkspaceResponse(ws *models.Workspace) Response {
	return Response{
		ID:        ws.ID,
		Name:      ws.Name,
		CreatedAt: ws.CreatedAt,
		UpdatedAt: ws.UpdatedAt,
	}
}
