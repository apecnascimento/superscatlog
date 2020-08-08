package api

import (
	"github.com/labstack/echo/v4"
)

// Routes return api pipeline
func Routes() (e *echo.Echo) {

	pipeline := echo.New()

	RegisterMiddlewares(pipeline)

	RegisterRoutes(pipeline)

	return pipeline

}
