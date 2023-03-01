package handler

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c echo.Context, statusCode int, message string) {

	logger, _ := zap.NewDevelopment()
	logger.Sugar().Error(message)

	c.JSON(statusCode, errorResponse{message})
}
