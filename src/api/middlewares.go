package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RegisterMiddlewares add middlewares to pipeline
func RegisterMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
}
