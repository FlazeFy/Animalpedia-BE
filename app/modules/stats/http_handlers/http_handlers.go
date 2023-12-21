package httphandlers

import (
	"app/modules/stats/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllTotalAnimalByZone(c echo.Context) error {
	ord := c.Param("ord")
	view := "animals_zone"

	result, err := repositories.GetAllTotalAnimalByZone("api/v1/stats/animalzone/"+ord, ord, view)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
