package routes

import (
	"net/http"

	lochandlers "tracker/modules/locations/http_handlers"
	stshandlers "tracker/modules/stats/http_handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitV1() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Tracker")
	})

	// =============== Public routes ===============

	// =============== Private routes (Admin) ===============
	e.POST("api/v1/location", lochandlers.PostLocation)
	e.GET("api/v1/location", lochandlers.GetAllLocations)
	e.GET("api/v2/location", lochandlers.GetAllLocationsFirebase)

	e.GET("api/v2/stats/location_category", stshandlers.GetTotalLocationByCategoryFirebase)
	e.GET("api/v2/stats/location_apps", stshandlers.GetTotalLocationByAppsFirebase)

	return e
}
