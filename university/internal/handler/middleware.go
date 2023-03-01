package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) UserIdentity(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		header := c.Request().Header.Get(authorizationHeader)

		if header == "" {
			NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		}

		headerParts := strings.Split(header, " ")

		if len(headerParts) != 2 {
			NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		}

		userId, err := h.svc.ParseToken(headerParts[1])

		if err != nil {
			NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		}

		c.Set(userCtx, userId)

		return next(c)

	}

}
