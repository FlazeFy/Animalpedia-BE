package httphandlers

import (
	middlewares "app/middlewares/jwt"
	"app/modules/auth/models"
	"app/modules/auth/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func PostLoginUser(c echo.Context) error {
	var body models.UserLogin
	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := middlewares.CheckLogin(c, body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostRegister(c echo.Context) error {
	var body models.UserRegister
	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := repositories.PostUserRegister(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SignOut(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Authorization token is missing"})
	}

	result, err := repositories.SignOut(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CheckRole(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Authorization token is missing"})
	}

	result, err := repositories.CheckRole(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
