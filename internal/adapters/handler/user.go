package handler

import (
	"net/http"
	core "trip-lot-api/internal/core/user"

	"trip-lot-api/util"

	"github.com/labstack/echo/v4"
)

type UserHTTPHandler struct {
	userService core.UserService
}

func NewUserHTTPHandler(userService core.UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
	}
}

func (h *UserHTTPHandler) CreateUser(c echo.Context) error {
	// Request body to struct
	var user core.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}

	// Create user
	_, err := h.userService.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "success")
}

func (h *UserHTTPHandler) GetUsers(c echo.Context) error {
	users, err := h.userService.GetUsers()
	if err != nil {
		return c.JSON(http.StatusOK, util.NewSuccessResponse(users))
	}
	return nil
}

func (h *UserHTTPHandler) GetUser(c echo.Context) error {
	userId := c.Param("id")

	user, err := h.userService.GetUser(userId)
	if err != nil {
		return c.JSON(http.StatusOK, user)
	}
	return nil
}

func (h *UserHTTPHandler) UpdateUser(c echo.Context) error {
	return nil
}

func (h *UserHTTPHandler) DeleteUser(c echo.Context) error {
	return nil
}
