package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	middleware2 "webreader/pkg/adapter/middleware"
)

// Path of route
const (
	QueryPath      = "/query"
	PlaygroundPath = "/playground"
)

// New creates route endpoint
func New(srv *handler.Server) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderXRequestedWith, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Use(middleware2.Auth())

	{
		e.POST(QueryPath, echo.WrapHandler(srv))
		e.GET(PlaygroundPath, func(c echo.Context) error {
			playground.Handler("GraphQL", QueryPath).ServeHTTP(c.Response(), c.Request())
			return nil
		})
	}

	return e
}
