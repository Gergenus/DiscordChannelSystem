package handler

import (
	"net/http"
	"strings"

	"github.com/Gergenus/pkg/tokens"
	"github.com/labstack/echo/v4"
)

type MiddlewareInterface interface {
	UserIndentity(ctx echo.Context) error
}

type EchoMiddleware struct {
	tokens tokens.TokenManager
}

func NewEchoMiddleware(tokens tokens.TokenManager) EchoMiddleware {
	return EchoMiddleware{tokens: tokens}
}

func (e *EchoMiddleware) UserIndentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		header := ctx.Request().Header.Get("Authorization")
		if header == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "empty auth header")
		}
		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid auth header")
		}

		token := headerParts[1]
		userId, err := e.tokens.ParseToken(token)
		if err != nil {
			ctx.Logger().Errorf("Failed to parse token: %v", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		ctx.Set("UserId", userId)
		return next(ctx)
	}
}
