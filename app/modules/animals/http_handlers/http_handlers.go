package httphandlers

import (
	"app/modules/animals/models"
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

func GetAnimalDetail(c echo.Context) error {
	slug := c.Param("slug")
	result, err := repositories.GetAnimalDetail(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAnimalCountryBySlug(c echo.Context) error {
	slug := c.Param("slug")
	result, err := repositories.GetAnimalCountryBySlug(slug)
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
	var obj models.PostAnimal

	// Data
	obj.AnimalName = c.FormValue("animals_name")
	obj.AnimalDesc = c.FormValue("animals_desc")
	obj.AnimalLatinName = c.FormValue("animals_latin_name")
	obj.AnimalImgUrl = c.FormValue("animals_img_url")
	obj.AnimalRegion = c.FormValue("animals_region")
	obj.AnimalZone = c.FormValue("animals_zone")
	obj.AnimalStatus = c.FormValue("animals_status")
	obj.AnimalCategory = c.FormValue("animals_category")

	result, err := repositories.PostAnimal(obj)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostNews(c echo.Context) error {
	var obj models.PostNews
	newsTimeReadStr := c.FormValue("news_time_read")

	// Converted form
	newsTimeReadInt, err := strconv.Atoi(newsTimeReadStr)
	if err != nil {
		newsTimeReadInt = 0
	}

	obj.NewsName = c.FormValue("news_name")
	obj.NewsTag = c.FormValue("news_tag")
	obj.NewsBody = c.FormValue("news_body")
	obj.NewsTimeRead = newsTimeReadInt
	obj.NewsImgUrl = c.FormValue("news_img_url")

	result, err := repositories.PostNews(obj)
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

func GetNewsDetail(c echo.Context) error {
	slug := c.Param("slug")
	result, err := repositories.GetNewsDetail(slug)
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

func RecoverAnimalBySlug(c echo.Context) error {
	slug := c.Param("slug")
	token := c.Request().Header.Get("Authorization")
	result, err := repositories.RecoverAnimalBySlug(slug, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func RecoverNewsBySlug(c echo.Context) error {
	slug := c.Param("slug")
	token := c.Request().Header.Get("Authorization")
	result, err := repositories.RecoverNewsBySlug(slug, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetNewsByTags(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	slug := c.Param("slug")
	result, err := repositories.GetNewsByTags(page, 10, "api/v1/news/slug/"+slug, slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetSourcesByContextSlug(c echo.Context) error {
	slug := c.Param("slug")
	types := c.Param("type")
	result, err := repositories.GetSourcesByContextSlug(types, slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelSourcesById(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.HardDelSourcesById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
