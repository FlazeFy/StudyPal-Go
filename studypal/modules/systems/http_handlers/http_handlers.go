package httphandlers

import (
	"net/http"
	"studypal/modules/systems/repositories"

	"github.com/labstack/echo"
)

func GetAllActiveDictionariesByType(c echo.Context) error {
	types := c.Param("type")
	result, err := repositories.GetDictionary("api/v1/dct/"+types, &types, true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetNotUsedDictionaries(c echo.Context) error {
	result, err := repositories.GetDictionary("api/v1/trash/dct", nil, false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
