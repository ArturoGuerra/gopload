package main

import (
    "handlers"
    "middleware"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/akrylysov/algnhsa"
)


func main () {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.POST("/upload", handlers.Upload, auth.UrlHeader, middleware.BodyLimit("20M"))
    //e.Logger.Fatal(e.Start(":5555"))
    opts := &algnhsa.Options{ BinaryContentTypes: []string{"image/jpeg", "image/png"}}
    algnhsa.ListenAndServe(e, opts)
}
