package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/arturoguerra/go-logging"
    "github.com/arturoguerra/goimgupload/internal/handlers"
    "github.com/arturoguerra/goimgupload/internal/mdlware"
    "github.com/arturoguerra/goimgupload/internal/config"
)

var log = logging.New()
var cfg = config.Load()

func init() {
    log.Infof("Upload Path: %s", cfg.Path)
}

func main () {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.POST(cfg.Path, handlers.Upload, mdlware.Auth, middleware.BodyLimit("100m"))
    e.GET("/", func(c echo.Context) error {
        return c.String(200, "")
    })

    e.Logger.Fatal(e.Start("0.0.0.0:5000"))
}
