package main

import (
    "handlers"
    "middleware"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
)

var Handle apigatewayproxy.Handler

func main () {
    e := echo.New()
    e.Listener = net.Listen()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.POST("/upload", handlers.Upload, auth.UrlHeader, middleware.BodyLimit("20M"))
    Handle = apigatewayproxy.New(e.Listener, nil).Handle
    //e.Logger.Fatal(e.Start(":5555"))
    go e.Start("")
}
