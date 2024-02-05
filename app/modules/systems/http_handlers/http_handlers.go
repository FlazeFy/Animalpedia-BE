package httphandlers

import (
	"app/modules/systems/models"
	"app/modules/systems/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetDictionaryByType(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	dctType := c.Param("type")
	result, err := repositories.GetDictionaryByType(page, 10, "api/v1/dct:"+dctType, dctType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelDictionaryById(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.HardDelDictionaryById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostDictionary(c echo.Context) error {
	result, err := repositories.PostDictionary(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostFeedback(c echo.Context) error {
	var obj models.PostFeedback
	fdbRateInt, _ := strconv.Atoi(c.FormValue("feedbacks_rate"))

	// Data
	obj.FdbRate = fdbRateInt
	obj.FdbDesc = c.FormValue("feedbacks_desc")

	result, err := repositories.PostFeedback(obj)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllFeedback(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("ord")
	ord_obj := c.Param("ord_obj")
	result, err := repositories.GetAllFeedback(page, 10, "api/v1/feedback/"+ord_obj+"/"+ord, ord_obj, ord)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
