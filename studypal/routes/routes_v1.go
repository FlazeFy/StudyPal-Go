package routes

import (
	"net/http"

	"github.com/labstack/echo"

	syshandlers "studypal/modules/systems/http_handlers"
)

func InitV1() *echo.Echo {
	e := echo.New()

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Studypal")
	})

	// =============== Public routes ===============

	// Dictionary
	e.GET("api/v1/dct/:type", syshandlers.GetAllActiveDictionariesByType)
	e.GET("api/v1/trash/dct", syshandlers.GetNotUsedDictionaries)

	// =============== Private routes (Admin) ===============

	return e
}
