package handlers

import (
    "fmt"
    "bytes"
    "strconv"
    "strings"
    "net/http"
    "mime/multipart"
    "github.com/sony/sonyflake"
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
    flake := sonyflake.NewSonyflake(sonyflake.Settings{})
    name,_ := flake.NextID()

    var buf bytes.Buffer
    buf.WriteString(strconv.FormatUint(name, 10))
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
    src,_ := file.Open()
    defer src.Close()
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"),
    }))
    uploader := s3manager.NewUploader(sess)
    _, err := uploader.Upload(&s3manager.UploadInput{
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

    src, err := file.Open()
    if err != nil {
        return err
    }

    defer src.Close()

    ext := GetExt(file.Filename)
    filename := GenFileName(ext)
    url := GenUrl(c, filename)
    err = S3Upload(file, filename)
    if err != nil {
        fmt.Println(err)
    }

    res := &Payload{filename, url}

    return c.JSON(http.StatusOK, res)
}
