package routes

import (
	"trip-lot-api/internal/adapters/handler"

	"github.com/labstack/echo/v4"
)

func NewUserRoute(e *echo.Echo, h *handler.UserHTTPHandler) {
	routes := e.Group("/api/users")

	routes.GET("", h.GetUsers)
	routes.GET("/:id", h.GetUser)
	routes.POST("", h.CreateUser)
	routes.PATCH("/:id", h.UpdateUser)
	routes.DELETE("/:id", h.DeleteUser)
}
