package auth

import (
	"app/configs"
	"app/packages/database"
	"fmt"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(pass string) string {
	hash, _ := HashPassword(pass)

	return hash
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetJWTConfiguration(name string) string {
	if name == "exp" {
		conf := configs.GetConfigJWT()
		return conf.JWT_EXP
	}
	return ""
}

func GetTokenHeader(c echo.Context) (bool, string) {
	authHeader := c.Request().Header.Get("Authorization")
	var userId string

	if authHeader == "" {
		return false, "No authorization header present"
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		return false, "Invalid authorization header format"
	}

	token := authHeader[len(bearerPrefix):]

	sqlStatement := `SELECT context_id FROM users_tokens WHERE token='` + token + `' LIMIT 1`

	fmt.Println(sqlStatement)

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return false, userId
	}

	// Map
	for rows.Next() {
		err = rows.Scan(&userId)
		if err != nil {
			return false, userId
		}
	}

	return true, userId
}
