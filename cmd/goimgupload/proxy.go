package main

import (
    "fmt"
    "net/url"
    "net/http"
    "net/http/httputil"
    "github.com/labstack/echo/v4"
)

func proxy(c echo.Context) error {
    url, err := url.Parse(cfg.ProxyURL)
    if err != nil {
        log.Error(err)
        return err
    }

    req := c.Request()
    res := c.Response().Writer

    p := httputil.NewSingleHostReverseProxy(url)
    req.URL.Host = url.Host
    req.URL.Scheme = url.Scheme
    req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
    req.Host = url.Host

    p.ErrorHandler = func(resp http.ResponseWriter, req *http.Request, err error) {
        c.Set("_error", echo.NewHTTPError(http.StatusBadGateway, fmt.Sprintf("Remove %s unreachable, could not forward: %v", url.String(), err)))
    }

    p.ServeHTTP(res, req)
    return nil
}
