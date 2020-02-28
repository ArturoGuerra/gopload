package main

import (
    "net/url"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/arturoguerra/go-logging"
    "github.com/arturoguerra/goimgupload/internal/handlers"
    "github.com/arturoguerra/goimgupload/internal/mdlware"
//    "github.com/akrylysov/algnhsa"
)

const rawurl = "arturo.minio.arturonet.com"
var log = logging.New()

func targets() []*middleware.ProxyTarget {
    url, err := url.Parse(rawurl)
    if err != nil {
        log.Fatal(err)
    }

    return []*middleware.ProxyTarget{
        {
            URL: url,
        },
    }
}

func main () {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.POST("/upload", handlers.Upload, mdlware.CallbackUrl, middleware.BodyLimit("20M"))
    e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets())))
    e.Logger.Fatal(e.Start(":5555"))
//    opts := &algnhsa.Options{ BinaryContentTypes: []string{"image/jpeg", "image/png"}}
//    algnhsa.ListenAndServe(e, opts)
}
