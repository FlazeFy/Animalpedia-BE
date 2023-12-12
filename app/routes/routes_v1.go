package routes

import (
	animalhandlers "app/modules/animals/http_handlers"
	syshandlers "app/modules/systems/http_handlers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitV1() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Animalpedia")
	})

	// =============== Public routes ===============

	// Dictionary
	e.GET("api/v1/dct/:type", syshandlers.GetDictionaryByType)

	// Animal
	e.GET("api/v1/animal/:ord", animalhandlers.GetAllAnimalHeaders)
	e.DELETE("api/v1/animal/by/:slug", animalhandlers.SoftDelAnimalBySlug)
	e.PUT("api/v1/animal/by/:slug", animalhandlers.UpdateAnimalBySlug)

	// News
	e.GET("api/v1/news/:ord", animalhandlers.GetAllNewsHeaders)
	e.DELETE("api/v1/news/by/:slug", animalhandlers.SoftDelNewsBySlug)

	// =============== Private routes ===============

	return e
}
