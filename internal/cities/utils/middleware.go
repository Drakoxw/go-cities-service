package utils

import (
	"github.com/labstack/echo/v4"
)

func ForceJSONMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		contentType := c.Request().Header.Get("Content-Type")
		if contentType != "application/json" {
			c.Request().Header.Set("Content-Type", "application/json")
		}
		return next(c)
	}
}
