package main

import (
    "handlers"
    "middleware"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)


func main () {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.POST("/upload", handlers.Upload, auth.UrlHeader, middleware.BodyLimit("20M"))

    e.Logger.Fatal(e.Start(":5555"))
}
