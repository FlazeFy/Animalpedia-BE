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
	e.DELETE("api/v1/animal/by/:slug", animalhandlers.SoftDelAnimalBySlug)
	e.DELETE("api/v1/animal/destroy/:slug", animalhandlers.HardDelAnimalBySlug)
	e.PUT("api/v1/animal/by/:slug", animalhandlers.UpdateAnimalBySlug)
	e.POST("api/v1/animal", animalhandlers.PostAnimal)

	// Tags
	e.GET("api/v1/tag/:ord", syshandlers.GetAllTags)
	e.DELETE("api/v1/tag/destroy/:id", syshandlers.HardDelTagById)
	e.POST("api/v1/tag", syshandlers.PostTag)

	// News
	e.GET("api/v1/news/:ord", animalhandlers.GetAllNewsHeaders)
	e.GET("api/v1/news/open/:slug", animalhandlers.GetNewsDetail)
	e.DELETE("api/v1/news/by/:slug", animalhandlers.SoftDelNewsBySlug)
	e.DELETE("api/v1/news/destroy/:slug", animalhandlers.HardDelNewsBySlug)
	e.PUT("api/v1/news/by/:slug", animalhandlers.UpdateNewsBySlug)
	e.POST("api/v1/news", animalhandlers.PostNews)

	// Comment
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

	// =============== Private routes ===============

	e.POST("api/v1/logout", authhandlers.SignOut, middlewares.CustomJWTAuth)
	e.GET("api/v1/check", authhandlers.CheckRole, middlewares.CustomJWTAuth)

	return e
}
