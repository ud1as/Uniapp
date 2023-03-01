package handler

import (
	"github.com/Studio56School/university/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) SignUp(c echo.Context) error {
	var input model.User

	if err := c.Bind(&input); err != nil {
		h.log.Sugar().Error(err)
	}

	id, err := h.svc.CreateUser(input)

	if err != nil {
		h.log.Sugar().Error(err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})

	return nil
}

type signInInput struct {
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

func (h *Handler) SignIn(c echo.Context) error {
	var input signInInput

	if err := c.Bind(&input); err != nil {
		h.log.Sugar().Error(err)
	}

	token, err := h.svc.GenerateToken(input.Username, input.Password)

	if err != nil {
		h.log.Sugar().Error(err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})

	return nil
}
