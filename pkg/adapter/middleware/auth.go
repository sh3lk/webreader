package middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"strings"
	"webreader/pkg/util/auth"
)

func extendRequestContext(c echo.Context, key any, val any) {
	c.SetRequest(
		c.Request().WithContext(
			context.WithValue(c.Request().Context(), key, val),
		),
	)
}

func AuthHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		authHeader := r.Header.Get(echo.HeaderAuthorization)
		if authHeader == "" {
			return next(c)
		}
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) == 2 && headerParts[0] == "Bearer" {
			auth, err := auth.ParseToken(headerParts[1])

			if err == nil {
				extendRequestContext(c, "auth", auth)
			}
		}

		return next(c)
	}
}

func Auth() func(next echo.HandlerFunc) echo.HandlerFunc {
	return AuthHandler
}
