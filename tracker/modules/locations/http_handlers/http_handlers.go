package httphandlers

import (
	"net/http"
	"strconv"
	"tracker/modules/locations/models"
	"tracker/modules/locations/repositories"

	"github.com/labstack/echo"
)

func GetAllLocations(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	result, err := repositories.GetAllLocations(page, 10, "api/v1/location")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllLocationsFirebase(c echo.Context) error {
	result, err := repositories.GetAllLocationsFirebase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostLocation(c echo.Context) error {
	var obj models.Location

	// Data
	obj.LocName = c.FormValue("location_name")
	obj.LocDesc = c.FormValue("location_desc")
	obj.LocLat = c.FormValue("location_lat")
	obj.LocLong = c.FormValue("location_long")
	obj.LocCat = c.FormValue("location_category")
	obj.LocAddr = c.FormValue("location_address")
	obj.LocApps = c.FormValue("location_apps")

	result, err := repositories.PostLocation(obj)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
