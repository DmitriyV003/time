package user_handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tracker/internal/application/dto"
	"tracker/internal/application/service"
)

type request struct {
	Name           string `json:"name" validate:"required"`
	Password       string `json:"password" validate:"required"`
	RepeatPassword string `json:"repeat_password" validate:"required,eqfield=Password"`
	Email          string `json:"email" validate:"required,email"`
}

type CreateHandler struct {
	userService *service.UserService
}

func (h *CreateHandler) handle(c echo.Context) error {
	var req request

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}

	par := dto.CreateUserDTO{
		Email:    req.Email,
		Name:     req.Name,
		LastName: "",
		Password: req.Password,
	}
	_, err = h.userService.Create(par)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error to register user")
	}

	return c.JSON(http.StatusCreated, nil)
}
