package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/arturoguerra/goimgupload/internal/handlers"
    "github.com/arturoguerra/goimgupload/internal/mdlware"
//    "github.com/akrylysov/algnhsa"
)


func main () {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.POST("/upload", handlers.Upload, mdlware.CallbackUrl, middleware.BodyLimit("20M"))
    e.Logger.Fatal(e.Start(":5555"))
//    opts := &algnhsa.Options{ BinaryContentTypes: []string{"image/jpeg", "image/png"}}
//    algnhsa.ListenAndServe(e, opts)
}
