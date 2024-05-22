package routes

import (
	"net/http"

	lochandlers "tracker/modules/locations/http_handlers"

	"github.com/labstack/echo"
)

func InitV1() *echo.Echo {
	e := echo.New()

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Tracker")
	})

	// =============== Public routes ===============

	// =============== Private routes (Admin) ===============
	e.POST("api/v1/location", lochandlers.PostLocation)
	e.GET("api/v1/location", lochandlers.GetAllLocations)
	e.GET("api/v2/location", lochandlers.GetAllLocationsFirebase)

	return e
}
