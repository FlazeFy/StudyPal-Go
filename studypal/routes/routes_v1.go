package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func InitV1() *echo.Echo {
	e := echo.New()

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Pet-Store")
	})

	// =============== Public routes ===============

	// Tag
	// e.GET("api/v1/tag", syshandlers.GetAllTag)

	// =============== Private routes (Admin) ===============

	return e
}