package httphandlers

import (
	"app/modules/comments/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllCommentBySlug(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	types := c.Param("type")
	slug := c.Param("slug")

	result, err := repositories.GetAllCommentBySlug(page, 10, "api/v1/comment/"+types+"/"+slug, types, slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

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
