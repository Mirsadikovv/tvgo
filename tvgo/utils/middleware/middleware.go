package middleware

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func XKeyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		xKeyStr := c.Request().Header.Get("X-Key")
		if xKeyStr == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "X-Key is missing")
		}

		xKey, err := strconv.ParseInt(xKeyStr, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "X-Key error")
		}

		if xKey%121 != 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "X-Key error")
		}

		return next(c)
	}
}
