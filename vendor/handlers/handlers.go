package handlers

import (
    "bytes"
    "strings"
    "net/http"
    "misc"
    "mime/multipart"
    "github.com/labstack/echo"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
)

type Config struct {
    Bucket string
}

var config Config = Config{"img-dixionary"}

func GetExt(f string) string {
    s := strings.Split(f, ".")
    return s[1]
}

func GenFileName(ext string) string {
    name := rand.RandString(10)

    var buf bytes.Buffer
    buf.WriteString(name)
    buf.WriteString(".")
    buf.WriteString(ext)
    result := buf.String()
    return result
}

func GenUrl(c echo.Context, filename string) string {
    var buf bytes.Buffer
    buf.WriteString(c.Request().Header.Get("url"))
    buf.WriteString("/")
    buf.WriteString(filename)
    url := buf.String()
    return url
}

func S3Upload(file *multipart.FileHeader, filename string) error {
    src, err := file.Open()
    if err != nil {
        return err
    }

    defer src.Close()
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"),
    }))
    uploader := s3manager.NewUploader(sess)
    _, err = uploader.Upload(&s3manager.UploadInput{
        Bucket: aws.String(config.Bucket),
        Key: aws.String(filename),
        Body: src,
        ACL: aws.String("public-read"),
        ContentType: aws.String("image/jpeg"),
    })



    return err


}

type Payload struct {
    Filename string `json:"filename"`
    Url string `json:"url"`
}

func Upload(c echo.Context) error {
    file, err := c.FormFile("file")

    if err != nil {
        return err
    }


    ext := GetExt(file.Filename)
    filename := GenFileName(ext)
    url := GenUrl(c, filename)
    err = S3Upload(file, filename)
    if err != nil {
        return err
    }

    res := &Payload{filename, url}

    return c.JSON(http.StatusOK, res)
}
