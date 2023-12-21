package httphandlers

import (
	"app/modules/stats/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllTotalAnimalByZone(c echo.Context) error {
	ord := c.Param("ord")
	view := "animals_zone"

	result, err := repositories.GetTotalStats("api/v1/stats/animalzone/"+ord, ord, view)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllTotalAnimalByStatus(c echo.Context) error {
	ord := c.Param("ord")
	view := "animals_status"

	result, err := repositories.GetTotalStats("api/v1/stats/animalstatus/"+ord, ord, view)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllTotalAnimalByCategory(c echo.Context) error {
	ord := c.Param("ord")
	view := "animals_category"

	result, err := repositories.GetTotalStats("api/v1/stats/animalstatus/"+ord, ord, view)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllTotalAnimalByRegion(c echo.Context) error {
	ord := c.Param("ord")
	view := "animals_region"

	result, err := repositories.GetTotalStats("api/v1/stats/animalstatus/"+ord, ord, view)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
