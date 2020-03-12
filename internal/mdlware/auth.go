package mdlware

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/goimgupload/internal/config"
)

var cfg = config.Load()

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        key := c.Request().Header.Get("X-API-KEY")
        if key == "" || key != cfg.ApiKey {
            return c.String(401, "Missing auth token")
        }

        url := c.Request().Header.Get("URL")
        if url != "" {
            return next(c)
        } else {
            return c.String(http.StatusBadRequest, "Invalid callback url")
        }
    }
}
