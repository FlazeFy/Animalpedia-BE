package httphandlers

import (
	"app/modules/animals/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllAnimalHeaders(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("ord")
	result, err := repositories.GetAllAnimalHeaders(page, 10, "api/v1/animal/"+ord, ord)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateAnimalBySlug(c echo.Context) error {
	slug := c.Param("slug")

	result, err := repositories.UpdateAnimalBySlug(slug, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostAnimal(c echo.Context) error {
	result, err := repositories.PostAnimal(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostNews(c echo.Context) error {
	result, err := repositories.PostNews(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllNewsHeaders(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("ord")
	result, err := repositories.GetAllNewsHeaders(page, 10, "api/v1/news/"+ord, ord)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateNewsBySlug(c echo.Context) error {
	slug := c.Param("slug")

	result, err := repositories.UpdateNewsBySlug(slug, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SoftDelAnimalBySlug(c echo.Context) error {
	slug := c.Param("slug")
	result, err := repositories.SoftDelAnimalBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelAnimalBySlug(c echo.Context) error {
	slug := c.Param("slug")
	result, err := repositories.HardDelAnimalBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelNewsBySlug(c echo.Context) error {
	slug := c.Param("slug")
	result, err := repositories.HardDelNewsBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SoftDelNewsBySlug(c echo.Context) error {
	slug := c.Param("slug")
	result, err := repositories.SoftDelNewsBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
