package httphandlers

import (
	"app/modules/comments/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func SoftDelCommentById(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.SoftDelCommentById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelCommentById(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.HardDelCommentById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCommentById(c echo.Context) error {
	id := c.Param("id")

	result, err := repositories.UpdateCommentById(id, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
