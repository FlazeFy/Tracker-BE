package httphandlers

import (
	"net/http"
	"tracker/modules/stats/repositories"

	"github.com/labstack/echo"
)

func GetTotalLocationByCategoryFirebase(c echo.Context) error {
	result, err := repositories.GetTotalLocationByCategoryFirebase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalLocationByAppsFirebase(c echo.Context) error {
	result, err := repositories.GetTotalLocationByAppsFirebase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
