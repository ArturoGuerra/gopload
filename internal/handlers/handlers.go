package handlers

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/go-logging"
    "github.com/arturoguerra/goimgupload/internal/config"
    "github.com/arturoguerra/goimgupload/internal/providers"
    "github.com/arturoguerra/goimgupload/internal/providers/utils"
)

var (
    cfg = config.Load()
    log = logging.New()
    provider providers.Uploader
)

func init() {
    p, err := providers.New(cfg)
    if err != nil {
        log.Fatal(err)
    }

    log.Info("Starting minio uploader")

    if err = p.Init(); err != nil {
        log.Fatal(err)
    }

    provider = p
}

func genUrl(c echo.Context, file string) string {
    return fmt.Sprintf("%s/%s", c.Request().Header.Get("url"), file)
}

func Upload(c echo.Context) error {
    file, err := c.FormFile("file")
    if err != nil {
        return err
    }

    object := utils.NewFile(file)
    err = provider.Upload(object)
    if err != nil {
        return err
    }

    url := genUrl(c, object.GetFilename())

    return c.JSON(http.StatusOK, map[string]string{
        "filename": object.GetFilename(),
        "url": url,
    })
}
