package middleware

import (
	"errors"
	"strings"

	"github.com/hfl0506/reader-app/internal/utils"
	"github.com/labstack/echo/v4"
)

func ValidateAuthHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")

		token := strings.Split(authorization, " ")[1]

		if token == "" {
			return errors.New("empty token string")
		}

		claims, err := utils.ValidateToken(token)

		if err != nil {
			return errors.New("claims in token failed")
		}

		c.Set("user", claims.Id)

		return next(c)
	})
}
