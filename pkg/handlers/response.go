package handlers

import (
	"github.com/labstack/echo/v4"
)

type errorResponse struct {
	Message string
	Status  int
}

func NewErrorResponse(ctx echo.Context, statusCode int, message string) error {
	return ctx.JSON(statusCode, errorResponse{Message: message, Status: statusCode})
}
