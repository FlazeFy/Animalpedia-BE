package routes

import (
	animalhandlers "app/modules/animals/http_handlers"
	comhandlers "app/modules/comments/http_handlers"
	stshandlers "app/modules/stats/http_handlers"
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
	e.DELETE("api/v1/dct/destroy/:id", syshandlers.HardDelDictionaryById)
	e.POST("api/v1/dct", syshandlers.PostDictionary)

	// Animal
	e.GET("api/v1/animal/:ord", animalhandlers.GetAllAnimalHeaders)
	e.DELETE("api/v1/animal/by/:slug", animalhandlers.SoftDelAnimalBySlug)
	e.DELETE("api/v1/animal/destroy/:slug", animalhandlers.HardDelAnimalBySlug)
	e.PUT("api/v1/animal/by/:slug", animalhandlers.UpdateAnimalBySlug)
	e.POST("api/v1/animal", animalhandlers.PostAnimal)

	// News
	e.GET("api/v1/news/:ord", animalhandlers.GetAllNewsHeaders)
	e.DELETE("api/v1/news/by/:slug", animalhandlers.SoftDelNewsBySlug)
	e.DELETE("api/v1/news/destroy/:slug", animalhandlers.HardDelNewsBySlug)
	e.PUT("api/v1/news/by/:slug", animalhandlers.UpdateNewsBySlug)
	e.POST("api/v1/news", animalhandlers.PostNews)

	// Comment
	e.DELETE("api/v1/comment/by/:id", comhandlers.SoftDelCommentById)
	e.DELETE("api/v1/comment/destroy/:id", comhandlers.HardDelCommentById)
	e.PUT("api/v1/comment/by/:slug", comhandlers.UpdateCommentById)

	// Stats
	e.GET("api/v1/stats/animalzone/:ord", stshandlers.GetAllTotalAnimalByZone)
	e.GET("api/v1/stats/animalstatus/:ord", stshandlers.GetAllTotalAnimalByStatus)
	e.GET("api/v1/stats/animalcategory/:ord", stshandlers.GetAllTotalAnimalByCategory)
	e.GET("api/v1/stats/animalregion/:ord", stshandlers.GetAllTotalAnimalByRegion)

	// =============== Private routes ===============

	return e
}
