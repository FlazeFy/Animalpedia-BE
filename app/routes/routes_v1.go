package routes

import (
	middlewares "app/middlewares/jwt"
	animalhandlers "app/modules/animals/http_handlers"
	authhandlers "app/modules/auth/http_handlers"
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

	// Auth
	e.POST("api/v1/login", authhandlers.PostLoginUser)
	e.POST("api/v1/register", authhandlers.PostRegister)

	// Dictionary
	e.GET("api/v1/dct/:type", syshandlers.GetDictionaryByType)
	e.DELETE("api/v1/dct/destroy/:id", syshandlers.HardDelDictionaryById)
	e.POST("api/v1/dct", syshandlers.PostDictionary)

	// Feedbacks
	e.POST("api/v1/feedbacks", syshandlers.PostFeedback)
	e.GET("api/v1/feedbacks/:ord_obj/:ord", syshandlers.GetAllFeedback)

	// Animal
	e.GET("api/v1/animal/:ord", animalhandlers.GetAllAnimalHeaders)
	e.GET("api/v1/animal/open/:slug", animalhandlers.GetAnimalDetail)
	e.GET("api/v1/animal/country/:slug", animalhandlers.GetAnimalCountryBySlug)

	e.DELETE("api/v1/animal/by/:slug", animalhandlers.SoftDelAnimalBySlug)
	e.DELETE("api/v1/animal/destroy/:slug", animalhandlers.HardDelAnimalBySlug)
	e.DELETE("api/v1/animal/country/destroy/:id", animalhandlers.HardDelAnimalCountryById)

	e.PUT("api/v1/animal/by/:slug", animalhandlers.UpdateAnimalBySlug)
	e.POST("api/v1/animal", animalhandlers.PostAnimal)
	e.POST("api/v1/animal/country", animalhandlers.PostAnimalCountry)
	e.POST("api/v1/animal/recover/:slug", animalhandlers.RecoverAnimalBySlug, middlewares.CustomJWTAuth)

	// Country
	e.GET("api/v1/country", animalhandlers.GetAllCountries)

	// Tags
	e.GET("api/v1/tag/:ord", syshandlers.GetAllTags)
	e.DELETE("api/v1/tag/destroy/:id", syshandlers.HardDelTagById)
	e.POST("api/v1/tag", syshandlers.PostTag)

	// News
	e.GET("api/v1/news/:ord", animalhandlers.GetAllNewsHeaders)
	e.GET("api/v1/news/open/:slug", animalhandlers.GetNewsDetail)
	e.GET("api/v1/news/tag/:slug", animalhandlers.GetNewsByTags)

	e.DELETE("api/v1/news/by/:slug", animalhandlers.SoftDelNewsBySlug)
	e.DELETE("api/v1/news/destroy/:slug", animalhandlers.HardDelNewsBySlug)
	e.PUT("api/v1/news/by/:slug", animalhandlers.UpdateNewsBySlug)
	e.POST("api/v1/news", animalhandlers.PostNews)
	e.POST("api/v1/news/recover/:slug", animalhandlers.RecoverNewsBySlug, middlewares.CustomJWTAuth)

	// Comment
	e.GET("api/v1/comment/:type/:slug", comhandlers.GetAllCommentBySlug)

	e.DELETE("api/v1/comment/by/:id", comhandlers.SoftDelCommentById)
	e.DELETE("api/v1/comment/destroy/:id", comhandlers.HardDelCommentById)
	e.PUT("api/v1/comment/by/:slug", comhandlers.UpdateCommentById)

	// Stats
	e.GET("api/v1/stats/animalzone/:ord", stshandlers.GetTotalAnimalByZone)
	e.GET("api/v1/stats/animalstatus/:ord", stshandlers.GetTotalAnimalByStatus)
	e.GET("api/v1/stats/animalcategory/:ord", stshandlers.GetTotalAnimalByCategory)
	e.GET("api/v1/stats/animalregion/:ord", stshandlers.GetTotalAnimalByRegion)
	e.GET("api/v1/stats/newstimeread/:ord", stshandlers.GetTotalNewsTimeRead)
	e.GET("api/v1/stats/commentcontext/:ord", stshandlers.GetTotalCommentContext)
	e.GET("api/v1/stats/animalpopcountry/:ord", stshandlers.GetTotalAnimalPopulationByCountry)
	e.GET("api/v1/stats/animalcountry/:ord", stshandlers.GetTotalAnimalVarietyByCountry)

	// Sources
	e.GET("api/v1/sources/:type/:slug", animalhandlers.GetSourcesByContextSlug)

	e.DELETE("api/v1/sources/:id", animalhandlers.HardDelSourcesById)

	// =============== Private routes ===============

	e.POST("api/v1/logout", authhandlers.SignOut, middlewares.CustomJWTAuth)
	e.GET("api/v1/check", authhandlers.CheckRole, middlewares.CustomJWTAuth)

	return e
}
