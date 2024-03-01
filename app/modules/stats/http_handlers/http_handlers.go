package httphandlers

import (
	"app/modules/stats/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func GetTotalAnimalByZone(c echo.Context) error {
	ord := c.Param("ord")
	view := "animals_zone"
	table := "animals"

	result, err := repositories.GetTotalStats(ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalAnimalByStatus(c echo.Context) error {
	ord := c.Param("ord")
	view := "animals_status"
	table := "animals"

	result, err := repositories.GetTotalStats(ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalAnimalByCategory(c echo.Context) error {
	ord := c.Param("ord")
	view := "animals_category"
	table := "animals"

	result, err := repositories.GetTotalStats(ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalAnimalByRegion(c echo.Context) error {
	ord := c.Param("ord")
	view := "animals_region"
	table := "animals"

	result, err := repositories.GetTotalStats(ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalNewsTimeRead(c echo.Context) error {
	ord := c.Param("ord")
	view := "news_time_read"
	table := "news"

	result, err := repositories.GetTotalStats(ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalCommentContext(c echo.Context) error {
	ord := c.Param("ord")
	view := "context_type"
	table := "comments"

	result, err := repositories.GetTotalStats(ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
